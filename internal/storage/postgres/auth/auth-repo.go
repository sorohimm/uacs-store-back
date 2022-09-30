package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"

	rbac "github.com/sorohimm/uacs-store-back/pkg/auth"
)

func NewAuthRepo(schema string, pool *pgxpool.Pool, salt string) *AuthRepo {
	return &AuthRepo{
		schema:   schema,
		pool:     pool,
		hashSalt: salt,
	}
}

type AuthRepo struct {
	schema   string
	pool     *pgxpool.Pool
	hashSalt string
}

func (o *AuthRepo) Registration(ctx context.Context, req *rbac.RegistrationRequest) (*User, error) {
	pwd := o.saltPassword(req.Password)
	var (
		tx  pgx.Tx
		err error
	)

	if tx, err = o.pool.BeginTx(ctx, pgx.TxOptions{}); err != nil {
		return nil, err
	}

	return saveUser(ctx, o.schema, tx, User{Username: req.Username, Email: req.Email, Password: pwd, Role: req.Role})
}

func (o *AuthRepo) saltPassword(password string) string {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(o.hashSalt))
	return fmt.Sprintf("%x", pwd.Sum(nil))
}

func (o *AuthRepo) Login(ctx context.Context, req *rbac.LoginRequest) error {

	return nil
}

func (o *AuthRepo) Logout(ctx context.Context, req *emptypb.Empty) error {
	return nil
}
