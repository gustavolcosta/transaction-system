package tests

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"transaction-system/domain/entities"
)

func TestNewAccount_ShouldWork(t *testing.T) {

	documentId := uuid.New().String()

	account, _ := entities.NewAccount(documentId)

	assert.Equal(t, documentId, account.DocumentId)
}

func TestNewAccount_When_DocumentIdIsEmpty_ShouldReturnAnError(t *testing.T) {

	documentId := ""
	errToReturn := errors.New("the document is required to create an account")

	account, err := entities.NewAccount(documentId)

	assert.Nil(t, account)
	assert.Equal(t, errToReturn, err)
}
