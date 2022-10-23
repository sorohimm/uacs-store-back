package cart

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
	"github.com/sorohimm/uacs-store-back/pkg/api"
	"github.com/sorohimm/uacs-store-back/pkg/log"
)

func NewRepo(schema string, pool *pgxpool.Pool) *Repo {
	return &Repo{
		schema: schema,
		pool:   pool,
	}
}

type Repo struct {
	schema string
	pool   *pgxpool.Pool
}

func (o *Repo) GetCart(ctx context.Context, req *api.CartReq) (*Cart, error) {
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

	var items []*Item
	if items, err = getCartItems(ctx, o.schema, tx, req.Id); err != nil {
		return nil, postgres.ResolveError(err)
	}
	cart := NewCart().SetID(req.Id).SetItems(items)

	return cart, nil
}

func getCartItems(ctx context.Context, schema string, tx pgx.Tx, cartID int64) ([]*Item, error) {
	sql := `
SELECT
id,
cart_id,
product_id,
quantity
FROM ` + schema + `.` + postgres.CartItemsTableName + `
WHERE cart_id=$1
`
	rows, err := tx.Query(ctx, sql, cartID)
	if err != nil {
		return nil, err
	}

	items := make([]*Item, 0, 0)
	for rows.Next() {
		var item *Item
		if err = rows.Scan(&item.ID, &item.CartID, &item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (o *Repo) DeleteCartItem(ctx context.Context, item *api.CartItem) error {
	var (
		tx     pgx.Tx
		err    error
		logger = log.FromContext(ctx).Sugar()
	)

	if tx, err = o.pool.BeginTx(ctx, pgx.TxOptions{}); err != nil {
		return err
	}
	defer func() {
		if err = postgres.CommitOrRollbackTx(ctx, tx, err); err != nil {
			logger.Errorf("tx: %s", err)
		}
	}()

	if err = deleteCartItem(ctx, o.schema, tx, item.Id); err != nil {
		return err
	}
	logger.Debugf("item deleted: %d", item.Id)

	return nil
}

func deleteCartItem(ctx context.Context, schema string, tx pgx.Tx, id int64) error {
	sql := `
DELETE FROM ` + schema + `.` + postgres.CartItemsTableName + `
WHERE id=$1;
`
	if _, err := tx.Exec(ctx, sql, id); err != nil {
		return err
	}

	return nil
}

func (o *Repo) AddCartItem(ctx context.Context, item *api.CartItem) (*Item, error) {
	var (
		id     int64
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

	iitem := NewItemFromApi(item)
	if id, err = addCartItem(ctx, o.schema, tx, iitem); err != nil {
		return nil, postgres.ResolveError(err)
	}
	iitem.SetID(id)

	return iitem, nil
}

func addCartItem(ctx context.Context, schema string, tx pgx.Tx, item *Item) (int64, error) {
	sql := `
INSERT INTO ` + schema + `.` + postgres.CartItemsTableName + `
(
cart_id,
product_id,
quantity
)
VALUES ($1,$2,$3)
RETURNING id;
`
	row := tx.QueryRow(ctx, sql, item.CartID, item.ProductID, item.Quantity)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
