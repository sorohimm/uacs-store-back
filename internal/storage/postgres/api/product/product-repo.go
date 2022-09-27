package product

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/shop/internal/log"
	"github.com/sorohimm/shop/internal/storage/postgres"
	"github.com/sorohimm/shop/pkg/api"
)

var ErrNotFound = errors.New("not found")

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
	sql := `SELECT 
		 id,
		 name,
		 price,
		 img,
		 FROM ` + o.schema + `.` + postgres.ProductTableName + ` WHERE id=$1`

	logger := log.FromContext(ctx).Sugar()
	logger.Debug(sql)

	row := o.pool.QueryRow(ctx, sql, id)

	var prod Product
	if err := row.Scan(
		&prod.Id,
		&prod.Name,
		&prod.Price,
		&prod.Img,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &prod, nil
}

func (o *ProductRepo) GetAllProducts(ctx context.Context, limit int64, offset int64) (Products, error) {
	return nil, nil
}

func (o *ProductRepo) CreateProduct(ctx context.Context, request *api.CreateProductRequest) (*Product, error) {
	return nil, nil
}
