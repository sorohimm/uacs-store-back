package product

import "github.com/jackc/pgx/v4/pgxpool"

func NewBrandRepo(schema string, pool *pgxpool.Pool) *InfoRepo {
	return &InfoRepo{
		pool:   pool,
		schema: schema,
	}
}

type BrandRepo struct {
	schema string
	pool   *pgxpool.Pool
}
