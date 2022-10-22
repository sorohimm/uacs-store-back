// Package category TODO
package category

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sorohimm/uacs-store-back/pkg/log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/pkg/product"
)

func NewCategoryRepo(schema string, pool *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		pool:   pool,
		schema: schema,
	}
}

type CategoryRepo struct {
	schema string
	pool   *pgxpool.Pool
}

func (o *CategoryRepo) CreateCategory(ctx context.Context, request *product.CreateCategoryRequest) (*Category, error) {
	var (
		id     int64
		tx     pgx.Tx
		err    error
		logger = log.FromContext(ctx).Sugar()
	)

	if tx, err = o.pool.BeginTx(ctx, pgx.TxOptions{}); err != nil {
		return nil, err
	}
	defer func() {
		if err = postgres.CommitOrRollbackTx(ctx, tx, err); err != nil {
			logger.Errorf("tx: %s", err)
		}
	}()

	category := NewCategoryFromRequest(request)

	if id, err = createCategory(ctx, o.schema, tx, category); err != nil {
		return nil, err
	}
	category.ID = id

	return category, nil
}

func createCategory(ctx context.Context, schema string, tx pgx.Tx, category *Category) (int64, error) {
	sql := `
INSERT INTO ` + schema + `.` + postgres.CategoryTableName + `
(
name
)
VALUES ($1)
RETURNING id;
`
	row := tx.QueryRow(ctx, sql, category.Name)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
