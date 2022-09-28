package product

import "github.com/jackc/pgx/v4/pgxpool"

func NewTypeRepo(schema string, pool *pgxpool.Pool) *InfoRepo {
	return &InfoRepo{
		pool:   pool,
		schema: schema,
	}
}

type TypeRepo struct {
	schema string
	pool   *pgxpool.Pool
}
