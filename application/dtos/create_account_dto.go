package dtos

type CreateAccountInputDTO struct {
	DocumentNumber string `json:"document_number"`
}

type CreateAccountOutputDTO struct {
	AccountId      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
