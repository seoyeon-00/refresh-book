package database

import (
	"fmt"
	"refresh-bookstore/configs"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func PGConnection(config configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	db.SetMaxOpenConns(config.DBMaxConnections)
	db.SetMaxIdleConns(config.DBMaxIdleConnections)
	db.SetConnMaxLifetime(time.Duration(config.DBMaxLifetimeConnections) * time.Second)

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
