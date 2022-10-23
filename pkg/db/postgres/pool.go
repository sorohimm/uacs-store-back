package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPGXPool(ctx context.Context, uri string, simpleProtocol bool) (*pgxpool.Pool, error) {
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
