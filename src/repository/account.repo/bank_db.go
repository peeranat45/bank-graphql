package bankrepo

import (
	"bank/src/pkg/data"
	"bank/src/pkg/data/models"
)


type bankAccountRepoDB struct {
	db *data.DB
}


func NewBankAccountRepo(db *data.DB) BankRepository {
	return &bankAccountRepoDB{db: db}
}

func (a bankAccountRepoDB) GetAccounts() ([]*models.BankAccount, error) {
	rows, err := a.db.Connection.Query("SELECT * FROM BankAccount")
	if err != nil {
		return nil ,err
	}

	var bankAccounts []*models.BankAccount

	for rows.Next() {
		var bankAccount models.BankAccount
		err := rows.Scan(&bankAccount.AccountID, &bankAccount.AccountNumber, &bankAccount.AccountHolderName, &bankAccount.AccountType, &bankAccount.Balance, &bankAccount.CreatedDate, &bankAccount.Status)
		if err != nil {
			return nil, err
		}
		bankAccounts = append(bankAccounts, &bankAccount)
	}

	return bankAccounts, nil
}

func (a bankAccountRepoDB) CreateAccount(b models.BankAccount) (*models.BankAccount, error) {
	query := "INSERT INTO BankAccount (AccountNumber, AccountHolderName, AccountType, Balance, CreatedDate, Status) VALUES (?, ?, ?, ?, ?, 'Active');"
	res, err := a.db.Connection.Exec(query, b.AccountNumber, b.AccountHolderName, b.AccountType, b.Balance, b.CreatedDate, b.Status)
	if err != nil {
		return nil ,err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	b.AccountID = int(id)
	return &b, nil
}

func (a bankAccountRepoDB) GetAccountByAccountNumber(accountNumber string) (*models.BankAccount, error) {
	query := "SELECT * FROM BankAccount WHERE BankAccount.AccountNumber=?"
	res:= a.db.Connection.QueryRow(query, accountNumber)
	var bankAccount models.BankAccount
	res.Scan(&bankAccount.AccountID, &bankAccount.AccountNumber, &bankAccount.AccountHolderName, &bankAccount.AccountType, &bankAccount.Balance, &bankAccount.CreatedDate, &bankAccount.Status)

	return &bankAccount, nil
}