package main

import (
	"context"

	"github.com/Shridhar2104/logilo/account"
	"github.com/Shridhar2104/logilo/graphql/models"
)

type mutationResolver struct {
	server *Server
}


func (r *mutationResolver) CreateAccount(ctx context.Context, input AccountInput) (*models.Account, error) {
	res, err := r.server.accountClient.CreateAccount(ctx, &account.Account{Name: input.Name})
	if err != nil {
		return nil, err
	}
	return &models.Account{ID: res.ID, Name: res.Name}, nil
}

