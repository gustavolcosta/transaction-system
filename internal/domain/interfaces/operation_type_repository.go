package interfaces

import (
	"transaction-system/internal/domain/entities"
)

type OperationTypeRepository interface {
	GetById(operationTypeId int) (*entities.OperationType, error)
}
