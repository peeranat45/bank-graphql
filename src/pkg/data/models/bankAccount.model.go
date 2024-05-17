package models

import (
	"time"
)

type BankAccount struct {
    AccountID        int       `json:"account_id"`
    AccountNumber    string    `json:"account_number"`
    AccountHolderName string   `json:"account_holder_name"`
    AccountType      string    `json:"account_type"`
    Balance          float64   `json:"balance"`
    CreatedDate      time.Time `json:"created_date"`
    Status           string    `json:"status"`
}
