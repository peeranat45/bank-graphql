package gql

import "github.com/graphql-go/graphql"


var CreateAccountInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateAccountInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"account_holder_name": {Type: graphql.NewNonNull(graphql.String)},
		"account_number":{Type: graphql.NewNonNull(graphql.String)},
		"account_type": {Type: graphql.NewNonNull(graphql.String)},
		"balance": {Type: graphql.NewNonNull(graphql.Float)},
		"status" : {Type: graphql.NewNonNull(graphql.String)},
	},
})

var CreateAccountArgument = graphql.FieldConfigArgument{
	"input": &graphql.ArgumentConfig{
		Type: CreateAccountInput,
	},
}

var CreateDepositInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateDepositInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"accound_id": {Type: graphql.NewNonNull(graphql.Int)},
		
	},
})
