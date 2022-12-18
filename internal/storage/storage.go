package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"recipe.core/internal/storage/psql"
)

type Storage struct {
	pgxConn *pgxpool.Pool
}

func New() (*Storage, error) {
	conn, err := psql.SetupDb()
	if err != nil {
		return nil, err
	}

	return &Storage{pgxConn: conn}, nil
}

func (s *Storage) GetPgxConn() *pgxpool.Pool {
	return s.pgxConn
}
