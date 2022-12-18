package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupDb() (*pgxpool.Pool, error) {
	connectUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		"postgres",
		"root",
		"localhost",
		5440,
		"recipe",
	)

	pool, err := pgxpool.New(context.Background(), connectUrl)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func Close(conn *pgxpool.Pool) {
	conn.Close()
}
