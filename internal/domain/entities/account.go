package entities

import (
	"errors"
	"strings"
)

type Account struct {
	Id             int
	DocumentNumber string
}

func NewAccount(documentNumber string) (*Account, error) {

	documentNumber = strings.TrimLeft(documentNumber, " ")

	if documentNumber == "" {
		return nil, errors.New("the document is required to create an account")
	}

	return &Account{
		DocumentNumber: documentNumber,
	}, nil
}
