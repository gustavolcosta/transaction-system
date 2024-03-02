package use_cases

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/application/use_cases"
	"transaction-system/internal/domain/entities"
	"transaction-system/internal/tests/implementations"
)

func TestCreateTransaction_ShouldWork(t *testing.T) {

	createTransactionUseCase := generateCreateTransaction()

	inputDTO := dtos.CreateTransactionInputDTO{
		AccountId:       1,
		OperationTypeId: 4,
		Amount:          123.40,
	}

	outputDTO, err := createTransactionUseCase.Execute(inputDTO)

	if err != nil {
		log.Printf("[TEST] Error to create transaction: %v", err)
	}

	assert.NotNil(t, outputDTO.TransactionId)
	assert.Equal(t, inputDTO.AccountId, outputDTO.AccountId)
	assert.Equal(t, inputDTO.OperationTypeId, outputDTO.OperationTypeId)
	assert.Equal(t, inputDTO.Amount, outputDTO.Amount)
}

func TestCreateTransaction_WhenAccountNotFound_ShouldReturnAnError(t *testing.T) {
	createTransactionUseCase := generateCreateTransaction()

	inputDTO := dtos.CreateTransactionInputDTO{
		AccountId:       99999,
		OperationTypeId: 4,
		Amount:          123.40,
	}

	assertError(createTransactionUseCase, t, inputDTO, "account not found")
}

func TestCreateTransaction_WhenOperationTypeNotFound_ShouldReturnAnError(t *testing.T) {
	createTransactionUseCase := generateCreateTransaction()

	inputDTO := dtos.CreateTransactionInputDTO{
		AccountId:       1,
		OperationTypeId: 9999,
		Amount:          123.40,
	}

	assertError(createTransactionUseCase, t, inputDTO, "operation type not found")
}

func assertError(createTransactionUseCase *use_cases.CreateTransactionUseCase, t *testing.T,
	inputDTO dtos.CreateTransactionInputDTO, expectedMessageError string) {

	errorToReturn := errors.New(expectedMessageError)

	outputDTO, err := createTransactionUseCase.Execute(inputDTO)

	assert.Nil(t, outputDTO)
	assert.Equal(t, errorToReturn, err)
}

func generateCreateTransaction() *use_cases.CreateTransactionUseCase {
	transactionRepository := implementations.NewTransactionRepositoryMemory()
	accountRepository := createAccountRepository()
	operationTypeRepository := implementations.NewOperationTypeRepositoryMemory()
	createTransactionUseCase := use_cases.NewCreateTransactionUseCase(transactionRepository, accountRepository, operationTypeRepository)

	return createTransactionUseCase
}

func createAccountRepository() *implementations.AccountRepositoryMemory {

	account := &entities.Account{
		Id:             1,
		DocumentNumber: uuid.New().String(),
	}

	accountRepository := implementations.NewAccountRepositoryMemory()

	err := accountRepository.Create(account)

	if err != nil {
		return nil
	}

	return accountRepository
}
