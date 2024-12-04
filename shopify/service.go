package shopify

import (
	"context"
	"errors"

	"github.com/Shridhar2104/logilo/shopify/pb"
)

type Service interface {
	CreateOrder(ctx context.Context, order *Order) error
	GetOrderByID(ctx context.Context, orderID string) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order) (*Order, error)
	FetchOrders(ctx context.Context, request *pb.FetchOrdersRequest) (*pb.FetchOrdersResponse, error)
	ListOrders(ctx context.Context, skip, take int) ([]*Order, error)
	SearchOrders(ctx context.Context, query string, skip, take int) ([]*Order, error)
	GetOrdersByIDs(ctx context.Context, orderIDs []string) ([]*Order, error)
}

type shopifyService struct {
	repo Repository
}

// CreateOrder implements Service.
func (s *shopifyService) CreateOrder(ctx context.Context, order *Order) error {
	return s.repo.CreateOrder(ctx, order)
}

// GetOrderByID implements Service.
func (s *shopifyService) GetOrderByID(ctx context.Context, orderID string) (*Order, error) {
	if orderID == "" {
		return nil, errors.New("orderID is required")
	}

	if order, err := s.repo.GetOrderByID(ctx, orderID); err != nil {
		return nil, err
	} else {
		return order, nil
	}

}
func (s *shopifyService) GetOrdersByIDs(ctx context.Context, orderIDs []string) ([]*Order, error) {
	return s.repo.GetOrdersByIDs(ctx, orderIDs)
}

// ListOrders implements Service.
func (s *shopifyService) ListOrders(ctx context.Context, skip int, take int) ([]*Order, error) {
	if take > 100 {
		take = 100
	}

	return s.repo.ListOrders(ctx, skip, take)
}
func (s *shopifyService) FetchOrders(ctx context.Context, request *pb.FetchOrdersRequest) (*pb.FetchOrdersResponse, error) {
	return s.repo.FetchOrders(ctx, request)
}

// SearchOrders implements Service.
func (s *shopifyService) SearchOrders(ctx context.Context, query string, skip int, take int) ([]*Order, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	return s.repo.SearchOrders(ctx, query, skip, take)
}

// UpdateOrder implements Service.
func (s *shopifyService) UpdateOrder(ctx context.Context, order *Order) (*Order, error) {
	return s.repo.UpdateOrder(ctx, order)
}

func NewService(repo Repository) Service {
	return &shopifyService{repo: repo}
}
