package main

type mutationResolver struct {
	server *Server
}

// func (r *mutationResolver) CreateAccount(ctx context.Context, input CreateAccountInput) (*Account, error) {
// 	return r.server.accountClient.CreateAccount(ctx, input)
// }

// func (r *mutationResolver) CreateOrder(ctx context.Context, input CreateOrderInput) (*Order, error) {
// 	return r.server.orderClient.CreateOrder(ctx, input)
// }

// func (r *mutationResolver) CreateShipment(ctx context.Context, input CreateShipmentInput) (*Shipment, error) {
// 	return r.server.shipmentClient.CreateShipment(ctx, input)
// }

// func (r *mutationResolver) CreateWalletTransaction(ctx context.Context, input CreateWalletTransactionInput) (*WalletTransaction, error) {
// 	return r.server.walletClient.CreateWalletTransaction(ctx, input)
// }
