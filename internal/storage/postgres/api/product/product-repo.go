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

func (o *ProductRepo) GetProductByID(ctx context.Context, id int64) (*Product, error) {
	sql := `
SELECT 
id,
name,
price
FROM ` + o.schema + `.` + postgres.ProductTableName + ` WHERE id=$1` // TODO: add image

	logger := log.FromContext(ctx).Sugar()
	logger.Debug(sql)

	row := o.pool.QueryRow(ctx, sql, id)

	var prod Product
	if err := row.Scan(
		&prod.ID,
		&prod.Name,
		&prod.Price,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &prod, nil
}

func (o *ProductRepo) GetAllProducts(ctx context.Context, limit int64, offset int64) (*Products, error) {
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

func (o *ProductRepo) GetAllProductsWithBrand(ctx context.Context, brandID int64, limit int64, offset int64) (*Products, error) {

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

func (o *ProductRepo) GetAllProductsWithType(ctx context.Context, typeID int64, limit int64, offset int64) (*Products, error) {
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

func (o *ProductRepo) GetAllProductsWithBrandAndType(ctx context.Context, typeID int64, brandID int64, limit int64, offset int64) (*Products, error) {
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

func (o *ProductRepo) scanAllProducts(rows pgx.Rows) (*Products, error) {
	var products Products
	for rows.Next() {
		var prod Product
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

func (o *ProductRepo) CreateProduct(ctx context.Context, request *api.CreateProductRequest) (*Product, error) {
	sql := `
INSERT INTO ` + o.schema + `.` + postgres.ProductTableName + `
(
name,
price,
brand_id,
type_id
)
VALUES  ($1,$2,$3,$4)
RETURNING id;` // Todo: add img field

	var (
		tx  pgx.Tx
		err error
	)

	if tx, err = o.pool.BeginTx(ctx, pgx.TxOptions{}); err != nil {
		return nil, err
	}
	defer commitOrRollbackTx(ctx, tx, err)

	row := tx.QueryRow(ctx, sql, request.Name, request.Price, request.BrandId, request.TypeId)

	var id int64
	if err = row.Scan(&id); err != nil {
		return nil, err
	}

	infoSql := `
INSERT INTO ` + o.schema + `.` + postgres.ProductInfoTableName + `
(
product_id,
title,
description
)
VALUES  ($1,$2,$3)
`
	for _, el := range request.Info {
		if _, err = tx.Exec(ctx, infoSql, id, el.Title, el.Description); err != nil {
			return nil, err
		}
	}

	product := NewProductFromRequest(request).SetID(id)

	return product, nil
}

func commitOrRollbackTx(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		return tx.Rollback(ctx)
	}
	return tx.Commit(ctx)
}
