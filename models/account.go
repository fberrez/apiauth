package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Account struct {
	ID          int    `json: "id"`
	Email       string `json: "email"`
	Password    string `json: "password"`
	ConnectedAt int    `json: "connectedAt"`
	CanceledAt  int    `json: "canceledAt"`
	Level       int    `json: "level"`
}

type AccountRepo struct {
	conn *pgx.Conn
}

func NewAccountRepo(conn *pgx.Conn) *AccountRepo {
	return &AccountRepo{
		conn: conn,
	}
}

// Create builds and executes a sql query to create a new account
func (u *AccountRepo) Create(ctx context.Context, email, password string, level int) error {
	tx, err := u.conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, "insert into accounts(email, password, level) values ($1, $2, $3)", email, password, level)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
