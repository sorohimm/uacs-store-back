// Package api TODO
package product

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
)

func addProductInfo(ctx context.Context, schema string, tx pgx.Tx, info []*dto.ProductInfo, productID int64) error {
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
