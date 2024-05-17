package gql

import (
	"bank/src/pkg/data/models"
	"bank/src/pkg/graphql/resolvers"

	"github.com/graphql-go/graphql"
)

type BankAccountQueries struct {
	// Getss	func() (*models.BankAccount, error)	`json:"get"`
}

type BankAccountMutation struct {
	CreateAccount func(map[string]interface{}) (*models.BankAccount, error) `json:"createAccount"`
}

type DepositTransactionQuries struct {

}


var BankAccountQueriesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BankAccountQueries",
	Fields: graphql.Fields{
		"gets": &graphql.Field{
			Type: graphql.NewList(BankAccountGraphQLType),
			Resolve: resolvers.GetAccountsResolver,
		},
	},
})

var BankAccountMutationsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BankAccountMutations",
	Fields: graphql.Fields{
		"createAccount": &graphql.Field{
			Type: BankAccountGraphQLType,
			Resolve: resolvers.CreateAccountResolver,
			Args: CreateAccountArgument,
		},
	},
})

var DepositTransactionQuriesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositTransactionQuries",
	
})



var BankAccountGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BankAccount",
	Fields: graphql.Fields{
		"account_id": &graphql.Field{Type: graphql.Int},
		"account_number": &graphql.Field{Type: graphql.String},
		"account_holder_name": &graphql.Field{Type: graphql.String},
		"balance": &graphql.Field{Type: graphql.Float},
		"created_date": &graphql.Field{Type: graphql.DateTime},
		"status": &graphql.Field{Type: graphql.String},
	},
})

var DepositeTransactionGrapQLType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Deposit",
	Fields: graphql.Fields{
		"transaction_id": &graphql.Field{Type: graphql.Int},
		"account_id": &graphql.Field{Type: graphql.Int},
		"transaction_date": &graphql.Field{Type: graphql.DateTime},
		"amount": &graphql.Field{Type: graphql.Float},
		"description": &graphql.Field{Type: graphql.String},

	},
})
	

