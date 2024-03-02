package main

import (
	"transaction-system/internal/infra/configs"
	"transaction-system/internal/infra/database"
)

func main() {

	configs.LoadEnv()
	db, err := database.ConnectDatabase()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	//log.Printf("Server started...")
	//err := http.ListenAndServe("localhost:8080", nil)
	//
	//if err != nil {
	//	log.Fatalf("Error to start server: %v", err)
	//}
}
