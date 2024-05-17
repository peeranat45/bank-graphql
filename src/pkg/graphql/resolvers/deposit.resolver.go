package resolvers

import (
	"bank/src/pkg/data"
	"bank/src/pkg/data/models"
	"bank/src/repository/deposit.repo"

	"github.com/graphql-go/graphql"
	"time"
)
func CreateDepositTransactionResolver(params graphql.ResolveParams) (interface{}, error) {
	db := data.NewDB()
	depositRepo := depositrepo.NewDepositRepo(db)

	createDeposit := models.DepositeTransaction {
		AccountID: 12,
		TransactionID: 12,
		TransactionDate: time.Now(),
		Amount: 1000,
		Description: "Deposit",

	}
	

	depositRepo.CreateDeposit(createDeposit)

	return createDeposit, nil

}