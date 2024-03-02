package main

import (
	"log"
	"transaction-system/internal/application/use_cases"
	"transaction-system/internal/infra/configs"
	"transaction-system/internal/infra/database"
	"transaction-system/internal/infra/database/adapters"
	"transaction-system/internal/infra/log_application"
)

func main() {

	configs.LoadEnv()
	db, err := database.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	accountRepository := adapters.NewAccountRepositoryPostgres(db)
	getAccount := use_cases.NewGetAccountByIdUseCase(accountRepository)

	outputDTO, err := getAccount.Execute(1)
	if err != nil {
		log_application.Error("Get account", err, "main")
	}

	log.Print(outputDTO)

	//log.Printf("Server started...")
	//err := http.ListenAndServe("localhost:8080", nil)
	//
	//if err != nil {
	//	log.Fatalf("Error to start server: %v", err)
	//}
}
