package main

import "context"

type accountResolver struct {
	server *Server

}

func (r *accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {
	// Extract account ID from the passed `obj` (Account).
	accountID := obj.ID

	// Call the orderClient to fetch the orders for the account.
	orders, err := r.server.shopifyClient.GetOrders(ctx, accountID)
	if err != nil {
		return nil, err // Return the error if the fetch fails.
	}

	return orders, nil // Return the fetched orders.
}