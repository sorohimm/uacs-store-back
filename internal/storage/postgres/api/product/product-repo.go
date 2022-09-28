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

func (o *ProductRepo) GetAllProducts(ctx context.Context, limit int64, offset int64) (*Products, error) {
	sql := `SELECT 
		 id,
		 name,
		 price,
		 img,
		 FROM ` + o.schema + `.` + postgres.ProductTableName + ` 
		 ORDER BY id
		 LIMIT $1 OFFSET $2;`

	logger := log.FromContext(ctx).Sugar()
	logger.Debug(sql)

	rows, err := o.pool.Query(ctx, sql, limit, offset)
	if err != nil {
		return nil, err
	}

	products, err := o.scanAllProducts(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (o *ProductRepo) GetAllProductsWithBrand(ctx context.Context, brandId int64, limit int64, offset int64) (*Products, error) {

	sql := `SELECT 
		 id,
		 name,
		 price,
		 img,
		 FROM ` + o.schema + `.` + postgres.ProductTableName + ` 
		 WHERE brand_id = $1
		 ORDER BY id
		 LIMIT $2 OFFSET $3;`

	logger := log.FromContext(ctx).Sugar()
	logger.Debug(sql)

	rows, err := o.pool.Query(ctx, sql, brandId, limit, offset)
	if err != nil {
		return nil, err
	}

	products, err := o.scanAllProducts(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (o *ProductRepo) GetAllProductsWithType(ctx context.Context, typeId int64, limit int64, offset int64) (*Products, error) {
	sql := `SELECT 
			id,
			name,
			price,
			img,
			FROM ` + o.schema + `.` + postgres.ProductTableName + ` 
			WHERE type_id = $1
			ORDER BY id
			LIMIT $2 OFFSET $3;`

	logger := log.FromContext(ctx).Sugar()
	logger.Debug(sql)

	rows, err := o.pool.Query(ctx, sql, typeId, limit, offset)
	if err != nil {
		return nil, err
	}

	products, err := o.scanAllProducts(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (o *ProductRepo) GetAllProductsWithBrandAndType(ctx context.Context, typeId int64, brandId int64, limit int64, offset int64) (*Products, error) {
	sql := `SELECT 
		 id,
		 name,
		 price,
		 img,
		 FROM ` + o.schema + `.` + postgres.ProductTableName + ` 
		 WHERE type_id = $1 AND brand_id = $2
		 ORDER BY id
		 LIMIT $3 OFFSET $4;`

	logger := log.FromContext(ctx).Sugar()
	logger.Debug(sql)

	rows, err := o.pool.Query(ctx, sql, typeId, brandId, limit, offset)
	if err != nil {
		return nil, err
	}

	products, err := o.scanAllProducts(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (o *ProductRepo) scanAllProducts(rows pgx.Rows) (*Products, error) {
	var products Products
	for rows.Next() {
		var prod Product
		if err := rows.Scan(
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
		products = append(products, &prod)
	}

	return &products, nil
}

func (o *ProductRepo) CreateProduct(ctx context.Context, request *api.CreateProductRequest) (*Product, error) {
	_ = `
INSERT INTO ` + o.schema + `.` + postgres.ProductTableName + `
(
name,
price,	
img,
brand_id,
type_id
)
VALUES  ($1,$2,$3,$4,$5)
ON CONFLICT (id) DO NOTHING`

	return nil, nil
}

func (o *ProductRepo) AddInfo(ctx context.Context) error {
	return nil
}
