package models

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// String is used to format the DB connection string
func (cfg DBConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func DefaultDBConfig() DBConfig {
	return DBConfig{
		Host:     "localhost",
		Port:     "5437",
		User:     "lenslock",
		Password: "password",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
}

func Open(config DBConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, fmt.Errorf("error opening db connection: %q", err)
	}

	return db, nil
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate set dialect error: %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("migrate up error: %w", err)
	}

	return nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	if dir == "" {
		dir = "."
	}
	
	goose.SetBaseFS(migrationsFS)

	defer func() {
		// remove FS in case there are other sections of app using goose and don't need FS
		goose.SetBaseFS(nil)
	}()
	
	return nil
}
