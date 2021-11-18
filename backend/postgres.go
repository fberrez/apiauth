package backend

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/juju/errors"
)

// NewPostgres initializes a new postgres connection.
func NewPostgres(ctx context.Context, url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, errors.Annotatef(err, "connect to postgres")
	}

	return conn, nil
}
