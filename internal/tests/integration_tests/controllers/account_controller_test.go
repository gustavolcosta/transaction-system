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
	"strconv"
	"testing"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/domain/exceptions"
	"transaction-system/internal/infra/configs"
	"transaction-system/internal/infra/database"
	"transaction-system/internal/infra/web/controllers"
	"transaction-system/internal/infra/web/response"
	"transaction-system/internal/tests/integration_tests/utils"
)

type AccountControllerTestSuite struct {
	suite.Suite
	Echo              *echo.Echo
	db                *sql.DB
	accountController *controllers.AccountController
}

func (suite *AccountControllerTestSuite) SetupTest() {

	suite.Echo = echo.New()

	db, err := database.ConnectDatabase()

	if err != nil {
		log.Fatalf("Error to connect database: %v", err)
	}

	suite.db = db

	suite.accountController = controllers.NewAccountController(db)
}

func (suite *AccountControllerTestSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *AccountControllerTestSuite) TestCreateAccount_ShouldWork() {

	payload := dtos.CreateAccountInputDTO{DocumentNumber: "14661471482"}

	payloadEncoded := utils.EncodePayload(payload)

	resp := suite.executeCreateAccount(payloadEncoded)

	var body dtos.CreateAccountOutputDTO

	err := json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		log.Fatalf("Erro create account controller test: %v", err)
	}

	assert.Equal(suite.T(), http.StatusCreated, resp.Code)
	assert.NotNil(suite.T(), body.AccountId)
	assert.Equal(suite.T(), payload.DocumentNumber, body.DocumentNumber)

}

func (suite *AccountControllerTestSuite) TestCreateAccount_WhenDocumentNumberIsEmpty_ShouldReturnBadRequest() {

	payload := dtos.CreateAccountInputDTO{DocumentNumber: ""}

	payloadEncoded := utils.EncodePayload(payload)

	resp := suite.executeCreateAccount(payloadEncoded)

	assert.Equal(suite.T(), http.StatusBadRequest, resp.Code)
}

func (suite *AccountControllerTestSuite) TestCreateAccount_WhenDocumentNumberIsBlank_ShouldReturnUnprocessableEntity() {

	payload := dtos.CreateAccountInputDTO{DocumentNumber: "  "}

	expectedResponse := exceptions.ValidationException{Message: "the document is required to create an account"}

	payloadEncoded := utils.EncodePayload(payload)

	resp := suite.executeCreateAccount(payloadEncoded)

	var body response.ExceptionResponse

	err := json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		log.Fatalf("Erro create account controller test: %v", err)
	}

	assert.Equal(suite.T(), http.StatusUnprocessableEntity, resp.Code)
	assert.Equal(suite.T(), body.Message, expectedResponse.Message)
}

func (suite *AccountControllerTestSuite) TestGetAccountById_ShouldWork() {

	//Create Account
	payload := dtos.CreateAccountInputDTO{DocumentNumber: "14661471482"}

	payloadEncoded := utils.EncodePayload(payload)

	respCreate := suite.executeCreateAccount(payloadEncoded)

	var createBody dtos.CreateAccountOutputDTO

	err := json.NewDecoder(respCreate.Body).Decode(&createBody)

	if err != nil {
		log.Fatalf("Erro create account controller test: %v", err)
	}
	//

	respGetAccount := suite.executeGetAccountById(strconv.Itoa(createBody.AccountId))

	var respBody dtos.GetAccountByIdOutputDTO

	err = json.NewDecoder(respGetAccount.Body).Decode(&respBody)

	if err != nil {
		log.Fatalf("Erro create account controller test: %v", err)
	}

	assert.Equal(suite.T(), http.StatusOK, respGetAccount.Code)
	assert.Equal(suite.T(), respBody.AccountId, createBody.AccountId)
	assert.Equal(suite.T(), respBody.DocumentNumber, createBody.DocumentNumber)
}

func (suite *AccountControllerTestSuite) TestGetAccountById_WhenIdIsNotANumber_ShouldReturnBadRequest() {
	respGetAccount := suite.executeGetAccountById("abc")

	suite.handleGetAccountByIdTestErrors(respGetAccount, "The accountId must be a number", http.StatusBadRequest)
}

func (suite *AccountControllerTestSuite) TestGetAccountById_WhenAccountNotFound_ShouldReturnNotFound() {
	respGetAccount := suite.executeGetAccountById("9999999")

	suite.handleGetAccountByIdTestErrors(respGetAccount, "account not found", http.StatusNotFound)
}

func (suite *AccountControllerTestSuite) handleGetAccountByIdTestErrors(respGetAccount *httptest.ResponseRecorder, messageError string, statusCode int) {
	var respBody response.ExceptionResponse

	err := json.NewDecoder(respGetAccount.Body).Decode(&respBody)

	if err != nil {
		log.Fatalf("Erro create account controller test: %v", err)
	}

	expectedResponse := response.ExceptionResponse{Message: messageError}

	assert.Equal(suite.T(), statusCode, respGetAccount.Code)
	assert.Equal(suite.T(), respBody.Message, expectedResponse.Message)
}

func (suite *AccountControllerTestSuite) executeCreateAccount(payload []byte) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	context := suite.Echo.NewContext(req, rec)
	err := suite.accountController.CreateAccount(context)

	if err != nil {
		log.Fatalf("Erro create account controller test: %v", err)
	}

	return rec
}

func (suite *AccountControllerTestSuite) executeGetAccountById(accountId string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	context := suite.Echo.NewContext(req, rec)
	context.SetPath("/accounts/:accountId")
	context.SetParamNames("accountId")
	context.SetParamValues(accountId)

	err := suite.accountController.GetAccountById(context)

	if err != nil {
		log.Fatalf("Erro create account controller test: %v", err)
	}

	return rec
}

func TestAccountControllerTestSuite(t *testing.T) {
	configs.LoadEnv("../../../../.env")
	suite.Run(t, new(AccountControllerTestSuite))
}
