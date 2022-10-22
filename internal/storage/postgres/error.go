package postgres

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var (
	ErrNotFound                              = errors.New("not found")
	ErrConflict                              = errors.New("conflict")
	DuplicateKeyViolatesUniqueConstraintCode = "23505"
)

func ResolveError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}

	var e *pgconn.PgError
	if errors.As(err, &e) {
		if err.(*pgconn.PgError).Code == DuplicateKeyViolatesUniqueConstraintCode {
			return ErrConflict
		}
	}

	return err
}
