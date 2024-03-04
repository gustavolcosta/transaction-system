package entities

import (
	"strings"
	"transaction-system/internal/domain/exceptions"
)

type Account struct {
	Id             int
	DocumentNumber string
}

func NewAccount(documentNumber string) (*Account, error) {

	documentNumber = strings.TrimLeft(documentNumber, " ")

	if documentNumber == "" {
		return nil, exceptions.NewValidationException("the document is required to create an account")
	}

	return &Account{
		DocumentNumber: documentNumber,
	}, nil
}
