package auth

import (
	"context"
	"github.com/sorohimm/uacs-store-back/pkg/log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
)

func NewAuthRepo(schema string, pool *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		schema: schema,
		pool:   pool,
	}
}

type AuthRepo struct {
	schema string
	pool   *pgxpool.Pool
}

func (o *AuthRepo) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
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

	user := &User{Username: req.User.Username, Email: req.User.Email, Password: req.User.Password, Role: req.User.Role}
	if user, err = saveUser(ctx, o.schema, tx, *user); err != nil {
		return nil, err
	}

	if err = saveSalt(ctx, o.schema, tx, user.ID, req.PwdSalt); err != nil {
		return nil, err
	}

	return user, nil
}

func (o *AuthRepo) GetUserByID(ctx context.Context, userID int64) (*User, error) {
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

	var user *User
	if user, err = getUserByID(ctx, o.schema, tx, userID); err != nil {
		return nil, err
	}

	return user, nil
}

func (o *AuthRepo) GetUserByUsername(ctx context.Context, username string) (*User, error) {
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

	var user *User
	if user, err = getUserByUsername(ctx, o.schema, tx, username); err != nil {
		return nil, err
	}

	return user, nil
}

func (o *AuthRepo) GetUserCredentialByUsername(ctx context.Context, username string) (*Credentials, error) {
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

	var credentials *Credentials
	if credentials, err = getCredentialsByUsername(ctx, o.schema, tx, username); err != nil {
		return nil, err
	}

	return credentials, nil
}
