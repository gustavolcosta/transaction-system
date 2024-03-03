package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"transaction-system/internal/domain/entities"
	"transaction-system/internal/domain/exceptions"
)

func TestNewAccount_ShouldWork(t *testing.T) {

	documentNumber := uuid.New().String()

	account, _ := entities.NewAccount(documentNumber)

	assert.Equal(t, documentNumber, account.DocumentNumber)
}

func TestNewAccount_When_DocumentIdIsEmpty_ShouldReturnAnError(t *testing.T) {

	documentId := ""
	errToReturn := exceptions.NewValidationException("the document is required to create an account")

	account, err := entities.NewAccount(documentId)

	assert.Nil(t, account)
	assert.Equal(t, errToReturn, err)
}
