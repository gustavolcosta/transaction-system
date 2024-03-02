package main

import (
	"log"
	"transaction-system/internal/application/dtos"
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
	opTypesRepository := adapters.NewOperationTypesRepositoryPostgres(db)
	transactionRepository := adapters.NewTransactionRepositoryPostgres(db)

	createTransaction := use_cases.NewCreateTransactionUseCase(transactionRepository, accountRepository, opTypesRepository)

	inputDTO := dtos.CreateTransactionInputDTO{
		AccountId:       1,
		OperationTypeId: 1,
		Amount:          -100,
	}

	outputDTO, err := createTransaction.Execute(inputDTO)

	if err != nil {
		log_application.Error("Error", err, "main")
		return
	}

	log.Print(outputDTO)

	//log.Printf("Server started...")
	//err := http.ListenAndServe("localhost:8080", nil)
	//
	//if err != nil {
	//	log.Fatalf("Error to start server: %v", err)
	//}
}
