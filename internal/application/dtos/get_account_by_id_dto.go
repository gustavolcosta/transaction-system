package dtos

type GetAccountByIdOutputDTO struct {
	AccountId      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
