package use_cases

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"transaction-system/internal/application/dtos"
	"transaction-system/internal/application/use_cases"
	"transaction-system/internal/tests/unit_tests/implementations"
)

func TestCreateAccount_ShouldWork(t *testing.T) {

	inputDTO := dtos.CreateAccountInputDTO{DocumentNumber: uuid.New().String()}
	accountRepository := implementations.NewAccountRepositoryMemory()
	createAccountUseCase := use_cases.NewCreateAccountUseCase(accountRepository)

	outputDTO, _ := createAccountUseCase.Execute(inputDTO)

	assert.NotNil(t, outputDTO.AccountId)
	assert.Equal(t, inputDTO.DocumentNumber, outputDTO.DocumentNumber)
}
