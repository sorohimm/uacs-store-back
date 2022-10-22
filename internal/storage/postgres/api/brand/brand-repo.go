// Package brand TODO
package brand

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sorohimm/uacs-store-back/pkg/log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/pkg/product"
)

func NewBrandRepo(schema string, pool *pgxpool.Pool) *BrandRepo {
	return &BrandRepo{
		pool:   pool,
		schema: schema,
	}
}

type BrandRepo struct {
	schema string
	pool   *pgxpool.Pool
}

func (o *BrandRepo) CreateBrand(ctx context.Context, request *product.CreateBrandRequest) (*Brand, error) {
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

	brand := NewBrandFromRequest(request)

	if id, err = createBrand(ctx, o.schema, tx, brand); err != nil {
		return nil, err
	}
	brand.ID = id

	return brand, nil
}

func createBrand(ctx context.Context, schema string, tx pgx.Tx, brand *Brand) (int64, error) {
	sql := `
INSERT INTO ` + schema + `.` + postgres.BrandTableName + `
(
name
)
VALUES ($1)
RETURNING id;
`
	row := tx.QueryRow(ctx, sql, brand.Name)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
