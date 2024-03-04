package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"os"
	_ "transaction-system/docs"
	"transaction-system/internal/infra/configs"
	"transaction-system/internal/infra/database"
	"transaction-system/internal/infra/web/route"
)

// @title Transaction System API
// @version 1.0
// @description App to handle with transactions routine
// @BasePath /
func main() {

	configs.LoadEnv(".env")

	db, err := database.ConnectDatabase()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	startServer(db)
}

func startServer(db *sql.DB) {
	port := os.Getenv("API_PORT")
	e := echo.New()

	route.Routes(e, db)

	e.Logger.Fatal(e.Start(":" + port))
}
