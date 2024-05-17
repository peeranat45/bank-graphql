package resolvers

import (
	"bank/src/pkg/data"
	"bank/src/pkg/data/models"
	"bank/src/repository/account.repo"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func validateBankAccount(input map[string]interface{}) map[string]interface{} {
	rules := map[string]interface{} {
		"account_number": "required",
		"account_holder_name": "required",
		"balance": "required",
		"status": "required",
	}
	// 	พังบรรทัดนี้
	errs := validate.ValidateMap(input, rules)
	fmt.Printf("Hello Wolrd\n")
	return errs
}

func GetAccountsResolver(params graphql.ResolveParams) (interface{}, error) {
	db := data.NewDB()
	accountRepo := bankrepo.NewBankAccountRepo(db)
	res, err := accountRepo.GetAccounts()
	if err != nil {
		return nil ,err
	}
	fmt.Println(res)
	return res, nil
}

func CreateAccountResolver(params graphql.ResolveParams) (interface{}, error) {
	
	db := data.NewDB()
	accountRepo := bankrepo.NewBankAccountRepo(db)
	// accountRepo.CreateAccount()
	

	input := params.Args["input"].(map[string]interface{})
	// invalid := validateBankAccount(input)

	// if invalid != nil{
	// 	return nil ,fmt.Errorf("%v", invalid) 
	// }
	
	
	newAccount := models.BankAccount{
		AccountNumber: input["account_number"].(string),
		AccountHolderName: input["account_holder_name"].(string),
		AccountType: input["account_holder_name"].(string),
		Balance: input["balance"].(float64),
		CreatedDate: time.Now(),
		Status: input["status"].(string),
	}
	

	res, err := accountRepo.CreateAccount(newAccount)
	if err != nil {
		return nil, err
	}
	return res, nil
}

