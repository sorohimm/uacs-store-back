package category

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/pkg/api"
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

func (o *CategoryRepo) CreateCategory(ctx context.Context, request *api.CreateCategoryRequest) (*Category, error) {
	sql := `
INSERT INTO ` + o.schema + postgres.CategoryTableName + `
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

	return &Category{ID: id, Name: request.Name}, nil
}
