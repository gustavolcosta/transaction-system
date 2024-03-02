package main

import (
	"github.com/google/uuid"
	"log"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/application/use_cases"
	"transaction-system/internal/infra/configs"
	"transaction-system/internal/infra/database"
	"transaction-system/internal/infra/database/adapters"
)

func main() {

	configs.LoadEnv()
	db, err := database.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	accountRepository := adapters.NewAccountRepositoryPostgres(db)
	createAccountUseCase := use_cases.NewCreateAccountUseCase(accountRepository)

	inputDTO := dtos.CreateAccountInputDTO{DocumentNumber: uuid.New().String()}

	outputDTO, err := createAccountUseCase.Execute(inputDTO)

	if err != nil {
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
