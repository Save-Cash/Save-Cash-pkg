package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitDB() error {
	dbURL := "postgres://username:password@localhost:5432/dbname"
	
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return fmt.Errorf("error parsing db config: %v", err)
	}

	dbPool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("failed to create db pool: %v", err)
	}

	return nil
}

func GetDB() *pgxpool.Pool {
	return dbPool
}

func CloseDB() {
	if dbPool != nil {
		dbPool.Close()
	}
}
