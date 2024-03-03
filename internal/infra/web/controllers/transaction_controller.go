package controllers

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/application/use_cases"
	"transaction-system/internal/domain/exceptions"
	"transaction-system/internal/infra/database/adapters"
	"transaction-system/internal/infra/web/response"
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
		return c.JSON(http.StatusBadRequest, response.NewExceptionResponse("Invalid payload"))
	}

	accountRepository := adapters.NewAccountRepositoryPostgres(transactionController.db)
	opTypeRepository := adapters.NewOperationTypesRepositoryPostgres(transactionController.db)
	transactionRepository := adapters.NewTransactionRepositoryPostgres(transactionController.db)
	createTransaction := use_cases.NewCreateTransactionUseCase(transactionRepository, accountRepository, opTypeRepository)

	outputDTO, err := createTransaction.Execute(inputDTO)

	if err != nil {

		var validationException *exceptions.ValidationException
		if errors.As(err, &validationException) {
			return c.JSON(http.StatusUnprocessableEntity, response.NewExceptionResponse(validationException.Error()))
		}

		var notFoundException *exceptions.NotFoundException
		if errors.As(err, &notFoundException) {
			return c.JSON(http.StatusNotFound, response.NewExceptionResponse(notFoundException.Error()))
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, outputDTO)
}
