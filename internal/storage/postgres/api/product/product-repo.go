package product

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/shop/pkg/api"
)

func NewProductRepo(pool *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		pool: pool,
	}
}

type ProductRepo struct {
	pool *pgxpool.Pool
}

func (o *ProductRepo) GetProductById(ctx context.Context, id int64) (*Product, error) {
	return nil, nil
}

func (o *ProductRepo) GetAllProducts(ctx context.Context) (Products, error) {
	return nil, nil
}

func (o *ProductRepo) CreateProduct(ctx context.Context, request *api.CreateProductRequest) (*Product, error) {
	return nil, nil
}
