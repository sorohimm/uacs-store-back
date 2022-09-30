package rbac

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
)

var ErrNotFound = errors.New("not found")

func getCredentials(ctx context.Context, schema string, tx pgx.Tx, userID int64) (*Credentials, error) {
	sql := `
SELECT email, password
FROM ` + schema + `.` + postgres.UserTableName + `
WHERE id=$1
`

	row := tx.QueryRow(ctx, sql, userID)

	var cred Credentials
	if err := row.Scan(
		&cred.Email,
		&cred.Password,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &cred, nil
}

func saveUser(ctx context.Context, schema string, tx pgx.Tx, user User) error {
	sql := `
INSERT INTO email, password ` + schema + `.` + postgres.UserTableName + `
(
user_id,
email,
password
role
)
VALUES  ($1,$2,$3,$4)
`

	if _, err := tx.Exec(ctx, sql, user.ID, user.Email, user.Password, user.Role); err != nil {
		return err
	}

	return nil
}
