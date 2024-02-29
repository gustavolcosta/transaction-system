package entities

import (
	"errors"
	"time"
)

type Transaction struct {
	Id            int
	AccountId     int
	OperationType *OperationType
	Amount        float64
	EventDate     time.Time
}

func NewTransaction(accountId int, operationType *OperationType, amount float64, eventDate time.Time) (*Transaction, error) {

	if amount <= 0 {
		return nil, errors.New("the amount of transaction must be greater than zero")
	}

	return &Transaction{
		AccountId:     accountId,
		Amount:        setAmountValue(amount, operationType.NegativeAmount),
		OperationType: operationType,
		EventDate:     eventDate,
	}, nil
}

func setAmountValue(amount float64, negativeAmount bool) float64 {

	if negativeAmount {
		return -amount
	}

	return amount
}
