package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/Shridhar2104/logilo/shopify"
)

type Server struct {
	// accountClient  *account.Client
	//shopifyClient    *shopify.Client
	// shipmentClient *shipment.Client
	// walletClient   *wallet.Client
}

func NewGraphQLServer(accountUrl, orderUrl, shipmentUrl, walletUrl string) (*Server, error) {

	// accountClient, err := account.NewClient(accountUrl)
	// if err != nil {
	// 	log.Fatalf("Failed to create account client: %v", err)
	// }

	// shopifyClient, err := shopify.NewClient(shopifyUrl)
	// if err != nil {
	// 	accountClient.Close()
	// 	log.Fatalf("Failed to create shopify client: %v", err)
	// }

	// shipmentClient, err := shipment.NewClient(shipmentUrl)
	// if err != nil {
	// 	accountClient.Close()
	// 	orderClient.Close()
	// 	log.Fatalf("Failed to create shipment client: %v", err)
	// }

	// walletClient, err := wallet.NewClient(walletUrl)
	// if err != nil {
	// 	accountClient.Close()
	// 	log.Fatalf("Failed to create wallet client: %v", err)
	// }

	return &Server{
		// accountClient: accountClient,
		// shopifyClient: shopifyClient,
		// shipmentClient: shipmentClient,
		// walletClient:   walletClient,
	}, nil
}


// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver {
// 	return &queryResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Account() AccountResolver {
// 	return &accountResolver{
// 		server: s,
// 	}
// }

func (s *Server) ToNewExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
