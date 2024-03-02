package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"transaction-system/internal/infra/log_application"
)

var contextLog = "DATABASE_CLIENT"

func ConnectDatabase() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log_application.Error("open connection with database", err, contextLog)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log_application.Error("connect with database", err, contextLog)
		return nil, err
	}

	log_application.Info("Connected to database", contextLog)

	return db, nil
}
