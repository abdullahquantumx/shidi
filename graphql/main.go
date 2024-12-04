package main

import (
	"log"
	"net/http"


	"github.com/99designs/gqlgen/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type AppConfig struct {
	AccountUrl string `env:"ACCOUNT_URL"`
	OrderUrl   string `env:"ORDER_URL"`
	ShipmentUrl string `env:"SHIPMENT_URL"`
	WalletUrl   string `env:"WALLET_URL"`
}

func main() {

	var config AppConfig
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}

	server, err := NewGraphQLServer(config.AccountUrl, config.OrderUrl, config.ShipmentUrl, config.WalletUrl)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	http.Handle("/graphql", handler.GraphQL(server.ToNewExecutableSchema()))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}