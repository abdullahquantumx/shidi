package main

import (
	"context"
	"time"

	"github.com/Shridhar2104/logilo/graphql/models"

)

type accountResolver struct {
	server *Server
}

func (r *accountResolver) Orders(ctx context.Context, obj *models.Account) ([]*Order, error) {

	res, err := r.server.shopifyClient.GetOrdersForShopAndAccount(ctx, obj.ShopName, obj.ID)
	if err != nil {
		return nil, err
		
	}
	orders := make([]*Order, len(res))
	for i, order := range res {
		orders[i] = &Order{ID: order.ID, Amount: order.TotalPrice, CreatedAt: order.CreatedAt.Format(time.RFC3339)}
	}
	return orders, nil
}
