package category

import "github.com/jackc/pgx/v4/pgxpool"

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
