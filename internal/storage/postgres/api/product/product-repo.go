package product

import (
	"context"
	"github.com/sorohimm/shop/internal/storage/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/shop/pkg/api"
)

func NewProductRepo(schema string, pool *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		pool:   pool,
		schema: schema,
	}
}

type ProductRepo struct {
	schema string
	pool   *pgxpool.Pool
}

func (o *ProductRepo) GetProductById(ctx context.Context, id int64) (*Product, error) {
	return nil, nil
}

func (o *ProductRepo) GetAllProducts(ctx context.Context) (Products, error) {
	return nil, nil
}

func (o *ProductRepo) CreateProduct(ctx context.Context, request *api.CreateProductRequest) (*Product, error) {
	_ = `SELECT 
			id, 
			entity_id,
			uri, 
			child, 
			parent, 
			ancestor,
			descendant, 
			entity,
			created, 
			loaded_in FROM ` + o.schema + `.` + postgres.ProductTableName + ` WHERE entity_id=$1`

	return nil, nil
}
