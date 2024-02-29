package entities

type OperationType struct {
	Id             int
	Description    string
	NegativeAmount bool
}

func NewOperationType(id int, description string, negativeAmount bool) *OperationType {
	return &OperationType{
		Id:             id,
		Description:    description,
		NegativeAmount: negativeAmount,
	}
}
