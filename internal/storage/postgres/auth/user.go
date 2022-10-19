// Package auth TODO
package auth

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
)

func getCredentialsByUsername(ctx context.Context, schema string, tx pgx.Tx, username string) (*Credentials, error) {
	var (
		user *User
		salt string
		err  error
	)

	if user, err = getUserByUsername(ctx, schema, tx, username); err != nil {
		return nil, err
	}
	if salt, err = getSalt(ctx, schema, tx, user.ID); err != nil {
		return nil, err
	}

	return &Credentials{PwdSalt: salt, UserID: user.ID, Email: user.Email, Username: user.Username, Password: user.Password}, nil
}

func saveSalt(ctx context.Context, schema string, tx pgx.Tx, userID int64, salt string) error {
	sql := `
INSERT INTO ` + schema + `.` + postgres.SaltTableName + `
(
user_id,
salt
)
VALUES ($1,$2)
`
	if _, err := tx.Exec(ctx, sql, userID, salt); err != nil {
		return err
	}
	return nil
}

func getSalt(ctx context.Context, schema string, tx pgx.Tx, userID int64) (string, error) {
	sql := `
SELECT salt FROM ` + schema + `.` + postgres.SaltTableName + `
WHERE user_id=$1
`
	row := tx.QueryRow(ctx, sql, userID)
	var salt string
	if err := row.Scan(&salt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrNotFound
		}
		return "", err
	}
	return salt, nil
}

func saveUser(ctx context.Context, schema string, tx pgx.Tx, user User) (*User, error) {
	sql := `
INSERT INTO ` + schema + `.` + postgres.UserTableName + `
(
username,
email,
password,
role
)
VALUES ($1,$2,$3,$4)
RETURNING id;
`
	row := tx.QueryRow(ctx, sql, user.Username, user.Email, user.Password, user.Role)

	var id int64
	if err := row.Scan(&id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		switch v := err.(type) {
		case *pgconn.PgError:
			if v.Code == "23505" {
				return nil, ErrUserAlreadyExists
			}
		default:
		}

		return nil, err
	}

	return &User{
			ID:       id,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role,
		},
		nil
}

func getUserByID(ctx context.Context, schema string, tx pgx.Tx, userID int64) (*User, error) {
	sql := `
SELECT
username,
email,
password,
role
FROM ` + schema + `.` + postgres.UserTableName + `
WHERE id=$1
`
	row := tx.QueryRow(ctx, sql, userID)
	user := User{ID: userID}
	if err := row.Scan(&user.Username, &user.Email, &user.Password, &user.Role); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func getUserByUsername(ctx context.Context, schema string, tx pgx.Tx, username string) (*User, error) {
	sql := `
SELECT
id,
email,
password,
role
FROM ` + schema + `.` + postgres.UserTableName + `
WHERE username=$1
`
	row := tx.QueryRow(ctx, sql, username)
	user := User{Username: username}
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}
