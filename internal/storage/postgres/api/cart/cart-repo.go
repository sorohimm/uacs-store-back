package cart

import (
	"context"
	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"
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

func (o *Repo) GetCart(ctx context.Context) (*Cart, error) {
	return nil, nil
}

func getCart(ctx context.Context, schema string, tx pgx.Tx) error {
	return nil
}

func (o *Repo) DeleteItemFromCart(ctx context.Context) error {
	return nil
}

func deleteItemFromCart(ctx context.Context, schema string, tx pgx.Tx) error {
	return nil
}

func (o *Repo) AddItemToCart(ctx context.Context) error {
	return nil
}

func addItemToCart(ctx context.Context, schema string, tx pgx.Tx) error {
	return nil
}
