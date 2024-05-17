package main

import (
	gql "bank/src/pkg/graphql"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)


func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: gql.RootQuery,
		Mutation: gql.RootMutation,
	})
	if err != nil {
		log.Fatal(err)
	}

	graphqlHanlder := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		Playground: true,
	})

	http.Handle("/bank", graphqlHanlder)
	http.ListenAndServe(":4000", nil)
}