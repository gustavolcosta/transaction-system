package interfaces

import "transaction-system/domain/entities"

type OperationTypeRepository interface {
	GetById(operationTypeId int) (*entities.OperationType, error)
}
