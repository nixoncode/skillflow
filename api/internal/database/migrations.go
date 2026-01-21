package database

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func setUpGoose() error {
	goose.SetBaseFS(migrationFiles)

	if err := goose.SetDialect("mysql"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)

	}

	return nil
}

func RunMigrations(db *sql.DB) error {
	if err := setUpGoose(); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func RollbackLastMigration(db *sql.DB) error {
	if err := setUpGoose(); err != nil {
		return err
	}

	if err := goose.Down(db, "migrations"); err != nil {
		return fmt.Errorf("failed to rollback last migration: %w", err)
	}

	return nil
}

func CreateMigration(name string) error {
	if err := setUpGoose(); err != nil {
		return err
	}

	if err := goose.Create(nil, "internal/database/migrations", name, "sql"); err != nil {
		return fmt.Errorf("failed to create migration: %w", err)
	}

	return nil
}

func RefreshMigrations(db *sql.DB) error {
	if err := setUpGoose(); err != nil {
		return err
	}

	if err := goose.Reset(db, "migrations"); err != nil {
		return fmt.Errorf("failed to reset migrations: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func MigrationStatus(db *sql.DB) error {
	if err := setUpGoose(); err != nil {
		return err
	}

	err := goose.Status(db, "migrations")
	if err != nil {
		return fmt.Errorf("failed to get migration status: %w", err)
	}

	return nil
}

func ResetMigrations(db *sql.DB) error {
	if err := setUpGoose(); err != nil {
		return err
	}

	if err := goose.Reset(db, "migrations"); err != nil {
		return fmt.Errorf("failed to reset migrations: %w", err)
	}

	return nil
}
