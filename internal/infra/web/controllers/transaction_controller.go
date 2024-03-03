package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/application/use_cases"
	"transaction-system/internal/infra/database/adapters"
)

type TransactionController struct {
	db *sql.DB
}

func NewTransactionController(db *sql.DB) *TransactionController {
	return &TransactionController{db: db}
}

func (transactionController *TransactionController) CreateTransaction(c echo.Context) error {

	var inputDTO dtos.CreateTransactionInputDTO

	if err := c.Bind(&inputDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Payload invalid",
		})
	}

	accountRepository := adapters.NewAccountRepositoryPostgres(transactionController.db)
	opTypeRepository := adapters.NewOperationTypesRepositoryPostgres(transactionController.db)
	transactionRepository := adapters.NewTransactionRepositoryPostgres(transactionController.db)
	createTransaction := use_cases.NewCreateTransactionUseCase(transactionRepository, accountRepository, opTypeRepository)

	outputDTO, err := createTransaction.Execute(inputDTO)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, outputDTO)
}
