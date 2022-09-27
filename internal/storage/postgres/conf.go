package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Config contains postgres configuration.
type Config struct {
	URI                   string `long:"uri" env:"URI" description:"PGX connection uri to the postgres" default:"postgresql://pg:test@localhost:5432/fer?sslmode=disable" required:"true"`
	DisableSimpleProtocol bool   `long:"simple.protocol"  description:"disable implicit prepared statement usage (if PreferSimpleProtocol == false wit bouncer usage it will produce errors in prepared statements)" `
}

func NewPGXPool(ctx context.Context, c *Config) (*pgxpool.Pool, error) {
	return newPGXPool(ctx, c.URI, !c.DisableSimpleProtocol)
}

func newPGXPool(ctx context.Context, uri string, simpleProtocol bool) (*pgxpool.Pool, error) {
	var (
		err    error
		conn   *pgxpool.Pool
		pgconf *pgxpool.Config
	)

	pgconf, err = pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, err
	}

	pgconf.ConnConfig.PreferSimpleProtocol = simpleProtocol

	if conn, err = pgxpool.ConnectConfig(ctx, pgconf); err != nil {
		return nil, fmt.Errorf("connect to postgres: %w", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return conn, nil
}
