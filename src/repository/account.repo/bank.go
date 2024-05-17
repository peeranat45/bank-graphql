package bankrepo

import "bank/src/pkg/data/models"

type BankRepository interface {
	GetAccounts() ([]*models.BankAccount, error)
	GetAccountByAccountNumber(string) (*models.BankAccount, error)
	CreateAccount(models.BankAccount) (*models.BankAccount, error)
}

