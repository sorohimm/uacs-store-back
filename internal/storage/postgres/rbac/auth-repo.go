package rbac

import (
	"context"
	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"

	rbac "github.com/sorohimm/uacs-store-back/pkg/rbac"
)

type AuthRepo struct {
	schema string
	pool   *pgxpool.Pool
}

func (o *AuthRepo) Registration(ctx context.Context, req *rbac.RegistrationRequest) error {
	var (
		tx  pgx.Tx
		err error
	)

	if tx, err = o.pool.BeginTx(ctx, pgx.TxOptions{}); err != nil {
		return err
	}

	return saveUser(ctx, o.schema, tx, User{Email: req.Email, Password: req.Password, Role: req.Role})
}

func (o *AuthRepo) Login(ctx context.Context, req *rbac.LoginRequest) error {
	return nil
}

func (o *AuthRepo) Logout(ctx context.Context, req *emptypb.Empty) error {
	return nil
}
