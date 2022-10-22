// Package product TODO
package product

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
	"github.com/sorohimm/uacs-store-back/pkg/api"
	"github.com/sorohimm/uacs-store-back/pkg/log"
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

func (o *ProductRepo) GetProductByID(ctx context.Context, id int64) (*dto.Product, error) {
	sql := `
SELECT 
id,
name,
price
FROM ` + o.schema + `.` + postgres.ProductTableName + ` WHERE id=$1` // TODO: add image

	var (
		tx     pgx.Tx
		err    error
		logger = log.FromContext(ctx).Sugar()
	)
	if tx, err = o.pool.BeginTx(ctx, pgx.TxOptions{}); err != nil {
		return nil, err
	}
	defer func() {
		if err = postgres.CommitOrRollbackTx(ctx, tx, err); err != nil {
			logger.Errorf("tx: %s", err)
		}
	}()

	row := tx.QueryRow(ctx, sql, id)

	var prod dto.Product
	if err = row.Scan(
		&prod.ID,
		&prod.Name,
		&prod.Price,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var info []*dto.ProductInfo
	if info, err = getInfo(ctx, o.schema, tx, prod.ID); err != nil {
		return nil, err
	}
	prod.Info = info

	return &prod, nil
}

func (o *ProductRepo) GetAllProducts(ctx context.Context, limit int64, offset int64) (*dto.Products, error) {
	sql := `
SELECT 
id,
name,
price,
img
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

func (o *ProductRepo) GetAllProductsWithBrand(ctx context.Context, brandID int64, limit int64, offset int64) (*dto.Products, error) {
	sql := `
SELECT 
id,
name,
price,
img
FROM ` + o.schema + `.` + postgres.ProductTableName + ` 
WHERE brand_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;`

	logger := log.FromContext(ctx).Sugar()
	logger.Debug(sql)

	rows, err := o.pool.Query(ctx, sql, brandID, limit, offset)
	if err != nil {
		return nil, err
	}

	products, err := o.scanAllProducts(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (o *ProductRepo) GetAllProductsWithType(ctx context.Context, typeID int64, limit int64, offset int64) (*dto.Products, error) {
	sql := `
SELECT 
id,
name,
price,
img
FROM ` + o.schema + `.` + postgres.ProductTableName + ` 
WHERE type_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;`

	logger := log.FromContext(ctx).Sugar()
	logger.Debug(sql)

	rows, err := o.pool.Query(ctx, sql, typeID, limit, offset)
	if err != nil {
		return nil, err
	}

	products, err := o.scanAllProducts(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (o *ProductRepo) GetAllProductsWithBrandAndType(ctx context.Context, typeID int64, brandID int64, limit int64, offset int64) (*dto.Products, error) {
	sql := `
SELECT 
id,
name,
price,
img
FROM ` + o.schema + `.` + postgres.ProductTableName + ` 
WHERE type_id = $1 AND brand_id = $2
ORDER BY id
LIMIT $3 OFFSET $4;`

	logger := log.FromContext(ctx).Sugar()
	logger.Debug(sql)

	rows, err := o.pool.Query(ctx, sql, typeID, brandID, limit, offset)
	if err != nil {
		return nil, err
	}

	products, err := o.scanAllProducts(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (o *ProductRepo) scanAllProducts(rows pgx.Rows) (*dto.Products, error) {
	var products dto.Products
	for rows.Next() {
		var prod dto.Product
		if err := rows.Scan(
			&prod.ID,
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

func (o *ProductRepo) CreateProduct(ctx context.Context, request *api.CreateProductRequest) (*dto.Product, error) {
	var (
		tx     pgx.Tx
		err    error
		logger = log.FromContext(ctx).Sugar()
	)

	if tx, err = o.pool.BeginTx(ctx, pgx.TxOptions{}); err != nil {
		return nil, err
	}
	defer func() {
		if err = postgres.CommitOrRollbackTx(ctx, tx, err); err != nil {
			logger.Errorf("tx: %s", err)
		}
	}()

	product := dto.NewProductFromRequest(request)
	id, err := createProduct(ctx, o.schema, tx, product)
	if err != nil {
		logger.Debugf("create api err: %s", err)
		return nil, err
	}
	product.SetID(id)

	if err = addProductInfo(ctx, o.schema, tx, product.Info, product.ID); err != nil {
		logger.Debugf("add api info err: %s", err)
		return nil, err
	}

	return product, nil
}

// createProduct inserts new api and returns id
func createProduct(ctx context.Context, schema string, tx pgx.Tx, product *dto.Product) (int64, error) {
	sql := `
INSERT INTO ` + schema + `.` + postgres.ProductTableName + `
(
name,
price,
brand_id,
type_id
)
VALUES  ($1,$2,$3,$4)
RETURNING id;` // Todo: add img field

	row := tx.QueryRow(ctx, sql, product.Name, product.Price, product.BrandID, product.TypeID)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
