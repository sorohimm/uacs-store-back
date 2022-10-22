///go:build integration

package product

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/uacs-store-back/pkg/db/postgres"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	docker "github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

var (
	tConn           *pgxpool.Pool
	testPostgresURI string
)

func TestProductRepo(t *testing.T) {
	initDatabase(t, testPostgresURI)

	dropDatabase(t, testPostgresURI)
}

const pathToSql = "../../../../scripts/migrate"

func initDatabase(t *testing.T, uri string) {
	t.Helper()
	m, err := migrate.New("file://"+pathToSql, uri)
	require.NoError(t, err)
	err = m.Up()
	if err != migrate.ErrNoChange {
		require.NoError(t, err)
	}
	serr, derr := m.Close()
	require.NoError(t, serr)
	require.NoError(t, derr)
}

func dropDatabase(t *testing.T, uri string) {
	t.Helper()
	m, err := migrate.New("file://"+pathToSql, uri)
	require.NoError(t, err)
	err = m.Down()
	if err != migrate.ErrNoChange {
		require.NoError(t, err)
	}
	serr, derr := m.Close()
	require.NoError(t, serr)
	require.NoError(t, derr)
}

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := docker.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	dbname := "y6t5r4e-test"
	env := []string{
		"TZ='GMT-3'",
		"PGTZ='GMT-3'",
		"POSTGRES_DB=" + dbname,
		"POSTGRES_USER=pg",
		"POSTGRES_PASSWORD=test",
	}
	resource, err := pool.Run("postgres", "latest", env)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		testPostgresURI = fmt.Sprintf("postgresql://pg:test@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), dbname)
		tConn, err = postgres.NewPGXPool(context.Background(), testPostgresURI, true)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
