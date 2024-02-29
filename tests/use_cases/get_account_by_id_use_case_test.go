package use_cases

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"transaction-system/application/dtos"
	"transaction-system/application/use_cases"
	"transaction-system/tests/implementations"
)

func TestGetAccountById_ShouldWork(t *testing.T) {

	accountRepository := implementations.NewAccountRepositoryMemory()
	inputCreateDTO := dtos.CreateAccountInputDTO{DocumentNumber: uuid.New().String()}
	createAccountUseCase := use_cases.NewCreateAccountUseCase(accountRepository)
	getAccountById := use_cases.NewGetAccountByIdUseCase(accountRepository)

	accountCreated, _ := createAccountUseCase.Execute(inputCreateDTO)

	account, _ := getAccountById.Execute(accountCreated.AccountId)

	assert.Equal(t, account.AccountId, accountCreated.AccountId)
	assert.Equal(t, account.DocumentNumber, accountCreated.DocumentNumber)
}

func TestGetAccountById_WhenAccountNotFound_ShouldReturnAnError(t *testing.T) {

	accountRepository := implementations.NewAccountRepositoryMemory()
	getAccountById := use_cases.NewGetAccountByIdUseCase(accountRepository)
	accountId := 99999

	errorToReturn := errors.New("account not found")

	account, err := getAccountById.Execute(accountId)

	assert.Nil(t, account)
	assert.Equal(t, errorToReturn, err)
}
