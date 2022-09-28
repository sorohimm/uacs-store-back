package product

import (
	"context"

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
	sql := `
INSERT INTO ` + o.schema + `.` + postgres.ProductInfoTableName + `
(
product_id,
title,
description
)
VALUES  ($1,$2,$3)
`
	for _, el := range info {
		if _, err := o.pool.Exec(ctx, sql, productID, el.Title, el.Description); err != nil {
			return err
		}
	}

	return nil
}
