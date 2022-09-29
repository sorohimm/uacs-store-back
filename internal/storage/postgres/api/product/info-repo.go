package product

import (
	"context"
	"github.com/sorohimm/shop/internal/storage/postgres/api/product/dto"

	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/shop/internal/storage/postgres"
	"github.com/sorohimm/shop/pkg/api"
)

func NewInfoRepo(schema string, pool *pgxpool.Pool) *InfoRepo {
	return &InfoRepo{
		pool:   pool,
		schema: schema,
	}
}

type InfoRepo struct {
	schema string
	pool   *pgxpool.Pool
}

func (o *InfoRepo) AddInfo(ctx context.Context, info []*api.ProductInfo, productID int64) error {
	var (
		tx  pgx.Tx
		err error
	)

	if tx, err = o.pool.BeginTx(ctx, pgx.TxOptions{}); err != nil {
		return err
	}
	defer postgres.CommitOrRollbackTx(ctx, tx, err)

	if err = addInfo(ctx, o.schema, tx, dto.NewProductInfosFromAPI(info), productID); err != nil {
		return err
	}

	return nil
}

func addInfo(ctx context.Context, schema string, tx pgx.Tx, info []*dto.ProductInfo, productID int64) error {
	sql := `
INSERT INTO ` + schema + `.` + postgres.ProductInfoTableName + `
(
product_id,
title,
description
)
VALUES  ($1,$2,$3)
`
	var err error
	for _, el := range info {
		if _, err = tx.Exec(ctx, sql, productID, el.Title, el.Description); err != nil {
			return err
		}
	}
	return nil
}

func (o *InfoRepo) GetInfo(ctx context.Context, productID string) (*dto.ProductInfo, error) {

	return nil, nil
}

func getInfo(ctx context.Context, schema string, tx pgx.Tx, productID int64) ([]*dto.ProductInfo, error) {
	sql := `
SELECT 
product_id,
title,
description
FROM ` + schema + `.` + postgres.ProductInfoTableName + ` WHERE product_id=$1`

	rows, err := tx.Query(ctx, sql, productID)
	if err != nil {
		return nil, err
	}

	var infos []*dto.ProductInfo
	for rows.Next() {
		var info dto.ProductInfo
		if err = rows.Scan(&info.ProductID, &info.Title, &info.Description); err != nil {
			return nil, err
		}
		infos = append(infos, &info)
	}

	return infos, nil
}
