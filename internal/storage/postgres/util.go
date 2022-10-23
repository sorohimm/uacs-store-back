// Package postgres TODO
package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func CommitOrRollbackTx(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		return tx.Rollback(ctx)
	}
	return tx.Commit(ctx)
}
