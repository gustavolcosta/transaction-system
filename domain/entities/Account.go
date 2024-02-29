package entities

import (
	"errors"
)

type Account struct {
	Id         int
	DocumentId string
}

func NewAccount(documentId string) (*Account, error) {

	if documentId == "" {
		return nil, errors.New("the document is required to create an account")
	}

	return &Account{
		DocumentId: documentId,
	}, nil
}
