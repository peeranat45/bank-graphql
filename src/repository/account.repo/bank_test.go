package bankrepo

import (
	"bank/src/pkg/data"
	"bank/src/pkg/data/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetBankRepo(t *testing.T) {
	t.Run("Get all account", func(t *testing.T) {
		//Arrange
		db := data.NewDB()
		bank := NewBankAccountRepo(db)

		//Act
		accounts, err := bank.GetAccounts()

		//Assert
		assert.Nil(t,err)
		assert.Greater(t, len(accounts), 2)
		
	}) 
	
	// t.Run("Get by accountNumber", func(t *testing.T) {
	// 	// Arrange
	// 	db := data.NewDB()
	// 	bank := NewBankAccountRepo(db)
	// 	ts, err :=time.Parse("2024-05-16 16:58:55", "2024-05-16 16:58:55")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	expetected := models.BankAccount {
	// 		AccountID: 1,
	// 		AccountNumber: "1234567890",
	// 		AccountType: "Savings",
	// 		Balance: 1000,
	// 		CreatedDate: ts,
	// 		Status: "Active",
	// 	}

	// 	// Act 
	// 	account, err := bank.GetAccountByAccountNumber("1234567890")

	// 	// Assert
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, expetected, account)

	// })
}

func TestCreateBankRepo(t *testing.T) {
	t.Run("Insert Bank Account", func(t *testing.T) {
		// Arrange
		db := data.NewDB()
		bank := NewBankAccountRepo(db)
		date := time.Now()
		newAccount := models.BankAccount{
			AccountNumber: "09877777",
			AccountHolderName: "Peanut",
			AccountType: "Savings",
			Balance: 0,
			CreatedDate: date,
			Status: "Active",
		}
		// Act
		acc, err := bank.CreateAccount(newAccount)
		// Assert
		assert.Nil(t, err)
		newAccount.AccountID = acc.AccountID
		assert.Equal(t, *acc, newAccount)
	})
}

