// Package brand TODO
package brand

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/pkg/api"
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

func (o *BrandRepo) CreateBrand(ctx context.Context, request *api.CreateBrandRequest) (*Brand, error) {
	sql := `
INSERT INTO ` + o.schema + `.` + postgres.BrandTableName + `
(
name
)
VALUES ($1)
RETURNING id;
`
	row := o.pool.QueryRow(ctx, sql, request.Name)

	var id int64
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &Brand{ID: id, Name: request.Name}, nil
}
