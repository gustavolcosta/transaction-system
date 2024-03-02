package main

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"transaction-system/internal/infra/configs"
	"transaction-system/internal/infra/database"
)

func main() {
	configs.LoadEnv()

	db, err := database.ConnectDatabase()
	defer db.Close()

	if err != nil {
		log.Fatalf("Error to connect database in migration: %v", err)
	}

	runMigration(db)
}

func runMigration(db *sql.DB) {

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/infra/database/migration",
		"postgres", driver)

	if err != nil {
		log.Printf("Error create migrate instance: %v", err)
		panic(err)
	}

	err = m.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Error to make migration: %v", err)
		panic(err)
	}

	log.Printf("Migration sucessfull")
}
