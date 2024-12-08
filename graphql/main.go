package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountUrl string `envconfig:"ACCOUNT_URL"`
	ShopifyUrl string `envconfig:"SHOPIFY_URL"`
}

func main() {

	var config AppConfig
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}

	server, err := NewGraphQLServer(config.AccountUrl, config.ShopifyUrl)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	http.Handle("/graphql", handler.GraphQL(server.ToNewExecutableSchema()))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	log.Fatal(http.ListenAndServe(":8084", nil))
}