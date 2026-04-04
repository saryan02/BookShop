package db

import (
	"database/sql"
	"fmt"

	"BookShop/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func PostgresDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return conn, nil
}
