package models

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg DBConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func DefaultDBConfig() DBConfig {
	return DBConfig{
		Host: "localhost",
		Port: "5437",
		User: "lenslock",
		Password: "password",
		Database: "lenslocked",
		SSLMode: "disable",
	}
}

func Open(config DBConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, fmt.Errorf("error opening db connection: %q", err)
	}

	return db, nil
}
