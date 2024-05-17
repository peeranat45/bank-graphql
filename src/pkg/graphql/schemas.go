package gql

import (
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"bankAccount": &graphql.Field{
				Type: BankAccountQueriesType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return &BankAccountQueries{}, nil
				},
			},
			"depositTransaction": &graphql.Field{
				Type: BankAccountMutationsType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return &DepositTransactionQuries{}, nil
				},
			},
		},
	},
)

var RootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation", 
		Fields: graphql.Fields{
			"bankAccount": &graphql.Field{
				Type: BankAccountMutationsType,
				Resolve:  func(p graphql.ResolveParams) (interface{}, error) {
					return &BankAccountMutation{}, nil
				},
			},
		},
	},
)