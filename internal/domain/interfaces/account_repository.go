package interfaces

import (
	"transaction-system/internal/domain/entities"
)

type AccountRepository interface {
	Create(account *entities.Account) error
	GetById(accountId int) (*entities.Account, error)
}
