package dto

import (
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId string `json:"customer_id"`
	AccountType string `json:"account_type"`
	Amount float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To Open a new account you need to deposite atleast 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}