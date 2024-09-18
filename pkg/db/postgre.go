package db

import (
	"database/sql"
	"fmt"
	"github.com/fichca/junior-astrologer-service/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func RunMigrations(dbConnection *sql.DB, cfg *config.Postgre) error {
	driver, err := postgres.WithInstance(dbConnection, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to get migration tool driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migrations",
		cfg.Name, driver)
	if err != nil {
		return fmt.Errorf("failed to connect migration tool: %w", err)
	}

	err = m.Up()
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func InitConnection(cfg *config.Postgre, logger *logrus.Logger) *sqlx.DB {
	db, err := sqlx.Connect(cfg.Driver, fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s port=%s", cfg.User,
		cfg.Name, cfg.SSLMode, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		logger.Fatal(err)
	}
	return db
}
