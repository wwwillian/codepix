package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID 			string 	`json:"id" validate:"required,uuid4"`
	AccountID 	string 	`json:"account_id" validate:"required,uuid4"`
	Amount 		float64	`json:"amount" validate:"required,numeric"`
	PixKeyTo 	string	`json:"pix_key_to" validate:"required"`
	PixKeyKind 	string	`json:"pix_key_kind" validate:"required"`
	Description string	`json:"description" validate:"required"`
	Status 		string	`json:"status" validate:"required"`
	Error 		string	`json:"error"`
}

func (transaction *Transaction) isValid() error {
	valid := validator.New()
	err := valid.Struct(transaction)
	if err != nil {
		fmt.Errorf("Error during transaction validation: %s", err.Error())
		return err
	}
	return nil
}

func (transaction *Transaction) ParseJson(data []byte) error {
	err := json.Unmarshal(data, transaction)
	if err != nil {
		return err
	}

	err = transaction.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (transaction *Transaction) ToJson() ([]byte, error) {
	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(transaction)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewTransaction() *Transaction  {
	return &Transaction{}
}