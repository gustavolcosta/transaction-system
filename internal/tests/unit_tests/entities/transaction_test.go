package entities

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	entities2 "transaction-system/internal/domain/entities"
)

func TestNewTransaction_ShouldWork(t *testing.T) {
	accountId := 1
	amount := 123.40
	eventDate := time.Now()
	operationType := entities2.NewOperationType(4, "PAGAMENTO", false)
	transaction, _ := entities2.NewTransaction(accountId, operationType, amount, eventDate)

	assert.Equal(t, accountId, transaction.AccountId)
	assert.Equal(t, operationType.Id, transaction.OperationType.Id)
	assert.Equal(t, amount, transaction.Amount)
	assert.Equal(t, eventDate, transaction.EventDate)
}

func TestNewTransaction_WhenAmountIsEqualToZero_ShouldReturnError(t *testing.T) {
	amount := 0
	accountId := 1
	eventDate := time.Now()
	operationType := entities2.NewOperationType(4, "PAGAMENTO", false)
	errorToReturn := errors.New("the amount of transaction must be greater than zero")

	transaction, err := entities2.NewTransaction(accountId, operationType, float64(amount), eventDate)

	assert.Nil(t, transaction)
	assert.Equal(t, errorToReturn, err)
}

func TestNewTransaction_WhenAmountIsLessThanZero_ShouldReturnError(t *testing.T) {
	amount := -1
	accountId := 1
	eventDate := time.Now()
	operationType := entities2.NewOperationType(4, "PAGAMENTO", false)
	errorToReturn := errors.New("the amount of transaction must be greater than zero")

	transaction, err := entities2.NewTransaction(accountId, operationType, float64(amount), eventDate)

	assert.Nil(t, transaction)
	assert.Equal(t, errorToReturn, err)
}

func TestNewTransaction_WhenOpTypeHasNegativeAmountTrue_ShouldHasNegativeAmount(t *testing.T) {

	operationType := entities2.NewOperationType(3, "SAQUE", true)
	amount := 123.40
	accountId := 1
	eventDate := time.Now()
	expectAmount := -amount

	transaction, _ := entities2.NewTransaction(accountId, operationType, amount, eventDate)

	assert.Equal(t, expectAmount, transaction.Amount)
}

func TestNewTransaction_WhenOpTypeHasNegativeAmountFalse_ShouldHasPositiveAmount(t *testing.T) {

	operationType := entities2.NewOperationType(4, "PAGAMENTO", false)
	amount := 123.40
	accountId := 1
	eventDate := time.Now()

	transaction, _ := entities2.NewTransaction(accountId, operationType, amount, eventDate)

	assert.Equal(t, amount, transaction.Amount)
}
