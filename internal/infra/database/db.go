package database

import (
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"strconv"
)

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
		log.Printf("Error to open connection with database %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error to connect with database %v", err)
		return nil, err
	}

	log.Println("Connected to database")

	return db, nil
}
