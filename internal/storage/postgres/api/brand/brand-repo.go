package brand

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
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

func (o *BrandRepo) CreateBrand(ctx context.Context) {}
