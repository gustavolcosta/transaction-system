package entities

import (
	"errors"
)

type Account struct {
	Id             int
	DocumentNumber string
}

func NewAccount(documentNumber string) (*Account, error) {

	if documentNumber == "" {
		return nil, errors.New("the document is required to create an account")
	}

	return &Account{
		DocumentNumber: documentNumber,
	}, nil
}
