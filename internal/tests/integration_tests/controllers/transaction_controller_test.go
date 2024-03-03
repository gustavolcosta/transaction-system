package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/domain/entities"
	"transaction-system/internal/infra/configs"
	"transaction-system/internal/infra/database"
	"transaction-system/internal/infra/database/adapters"
	"transaction-system/internal/infra/web/controllers"
	"transaction-system/internal/infra/web/response"
	"transaction-system/internal/tests/integration_tests/utils"
)

type TransactionControllerTestSuite struct {
	suite.Suite
	Echo                  *echo.Echo
	db                    *sql.DB
	transactionController *controllers.TransactionController
	accountId             int
}

func (suite *TransactionControllerTestSuite) SetupTest() {

	suite.Echo = echo.New()

	db, err := database.ConnectDatabase()

	if err != nil {
		log.Fatalf("Error to connect database: %v", err)
	}

	suite.db = db

	suite.transactionController = controllers.NewTransactionController(db)

	suite.accountId = createAccount(db)
}

func (suite *TransactionControllerTestSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *TransactionControllerTestSuite) TestCreateTransaction_ShouldWork() {

	payload := dtos.CreateTransactionInputDTO{
		AccountId:       suite.accountId,
		OperationTypeId: 1,
		Amount:          100,
	}

	payloadEncoded := utils.EncodePayload(payload)

	resp := suite.executeCreateTransaction(payloadEncoded)

	var respBody dtos.CreateTransactionOutputDTO

	err := json.NewDecoder(resp.Body).Decode(&respBody)

	if err != nil {
		log.Fatalf("Erro create transaction controller test: %v", err)
	}

	assert.Equal(suite.T(), http.StatusCreated, resp.Code)
	assert.NotNil(suite.T(), respBody.TransactionId)
	assert.Equal(suite.T(), payload.AccountId, respBody.AccountId)
	assert.Equal(suite.T(), payload.OperationTypeId, respBody.OperationTypeId)
	assert.Equal(suite.T(), payload.Amount, respBody.Amount)
}

func (suite *TransactionControllerTestSuite) TestCreateTransaction_WhenAccountNotFound_ShouldReturnNotFound() {
	payload := dtos.CreateTransactionInputDTO{
		AccountId:       999999,
		OperationTypeId: 1,
		Amount:          100,
	}

	suite.handleTestsException(payload, "account not found", http.StatusNotFound)
}

func (suite *TransactionControllerTestSuite) TestCreateTransaction_WhenOpTypeNotFound_ShouldReturnNotFound() {
	payload := dtos.CreateTransactionInputDTO{
		AccountId:       suite.accountId,
		OperationTypeId: 99999,
		Amount:          100,
	}

	suite.handleTestsException(payload, "operation type not found", http.StatusNotFound)
}

func (suite *TransactionControllerTestSuite) TestCreateTransaction_WhenAmountIsNegative_ShouldReturnAnError() {
	payload := dtos.CreateTransactionInputDTO{
		AccountId:       suite.accountId,
		OperationTypeId: 4,
		Amount:          -100,
	}

	suite.handleTestsException(payload, "the amount of transaction must be greater than zero", http.StatusUnprocessableEntity)
}

func (suite *TransactionControllerTestSuite) handleTestsException(payload dtos.CreateTransactionInputDTO, messageException string, statusCode int) {
	payloadEncoded := utils.EncodePayload(payload)

	resp := suite.executeCreateTransaction(payloadEncoded)

	var respBody response.ExceptionResponse

	expectedException := response.ExceptionResponse{Message: messageException}

	err := json.NewDecoder(resp.Body).Decode(&respBody)

	if err != nil {
		log.Fatalf("Erro create transaction controller test: %v", err)
	}

	assert.Equal(suite.T(), statusCode, resp.Code)
	assert.Equal(suite.T(), expectedException.Message, respBody.Message)
}

func (suite *TransactionControllerTestSuite) executeCreateTransaction(payload []byte) *httptest.ResponseRecorder {

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	context := suite.Echo.NewContext(req, rec)
	err := suite.transactionController.CreateTransaction(context)

	if err != nil {
		log.Fatalf("Erro create transaction controller test: %v", err)
	}

	return rec
}

func createAccount(db *sql.DB) int {

	account := entities.Account{
		DocumentNumber: "05560550043",
	}

	accountRepository := adapters.NewAccountRepositoryPostgres(db)
	err := accountRepository.Create(&account)

	if err != nil {
		log.Fatalf("Error to create account in transaction controller test: %v", err)
	}

	return account.Id
}

func TestTransactionControllerTestSuite(t *testing.T) {
	configs.LoadEnv("../../../../.env")
	suite.Run(t, new(TransactionControllerTestSuite))
}
