package controllers

import (
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/application/use_cases"
	"transaction-system/internal/domain/exceptions"
	"transaction-system/internal/infra/database/adapters"
	"transaction-system/internal/infra/log_application"
	"transaction-system/internal/infra/web/response"
)

type AccountController struct {
	db *sql.DB
}

func NewAccountController(db *sql.DB) *AccountController {
	return &AccountController{db: db}
}

func (accountController *AccountController) CreateAccount(c echo.Context) error {
	var inputDTO dtos.CreateAccountInputDTO

	if err := c.Bind(&inputDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewExceptionResponse("Invalid payload"))
	}

	validate := validator.New()
	if err := validate.Struct(inputDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewExceptionResponse("document_number is required"))
	}

	accountRepository := adapters.NewAccountRepositoryPostgres(accountController.db)
	createAccountUseCase := use_cases.NewCreateAccountUseCase(accountRepository)

	outputDTO, err := createAccountUseCase.Execute(inputDTO)

	if err != nil {

		var validationException *exceptions.ValidationException
		if errors.As(err, &validationException) {
			return c.JSON(http.StatusUnprocessableEntity, response.NewExceptionResponse(validationException.Error()))
		}

		return c.JSON(http.StatusInternalServerError, response.NewExceptionResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, outputDTO)
}

func (accountController *AccountController) GetAccountById(c echo.Context) error {

	accountId, err := strconv.Atoi(c.Param("accountId"))

	if err != nil {
		log_application.Error("Convert accountId to int", err, "GET_ACCOUNT_BY_ID_CONTROLLER")
		return c.JSON(http.StatusBadRequest, response.NewExceptionResponse("The accountId must be a number"))
	}

	accountRepository := adapters.NewAccountRepositoryPostgres(accountController.db)
	getAccountById := use_cases.NewGetAccountByIdUseCase(accountRepository)

	outputDTO, err := getAccountById.Execute(accountId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewExceptionResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, outputDTO)
}
