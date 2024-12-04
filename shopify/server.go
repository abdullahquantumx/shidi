package shopify

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/Shridhar2104/logilo/shopify/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedShopifyServiceServer
	service Service
}

// NewGRPCServer initializes and starts a new gRPC server
func NewGRPCServer(service Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", port, err)
	}

	server := grpc.NewServer()
	pb.RegisterShopifyServiceServer(server, &grpcServer{
		service: service,
	})
	reflection.Register(server)

	fmt.Printf("gRPC server started on port %d\n", port)
	return server.Serve(lis)
}

// CreateOrder handles the gRPC call to create an order
func (s *grpcServer) CreateOrder(ctx context.Context, r *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := &Order{
		ID:        r.Order.Id,
		AccountID: r.Order.AccountID,
		Phase:     r.Order.Phase,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.service.CreateOrder(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	return &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id: order.ID,
		},
	}, nil
}

// GetOrderByID handles the gRPC call to get an order by its ID
func (s *grpcServer) GetOrderByID(ctx context.Context, r *pb.GetOrderByIDRequest) (*pb.GetOrderByIDResponse, error) {
	order, err := s.service.GetOrderByID(ctx, r.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order by ID: %w", err)
	}

	return &pb.GetOrderByIDResponse{
		Order: &pb.Order{
			Id:        order.ID,
			AccountID: order.AccountID,
			CreatedAt: order.CreatedAt.Format(time.RFC3339),
			UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
			Phase:     order.Phase,
		},
	}, nil
}

// ListOrders handles the gRPC call to list orders
func (s *grpcServer) ListOrders(ctx context.Context, r *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {

	var res []*Order
	var err error
	if r.Query != "" {
		res, err = s.service.SearchOrders(ctx, r.Query, r.Skip, r.Take)
	} else if r.Ids != nil {
		res, err = s.service.GetOrdersByIDs(ctx, r.Ids)
	} else {
		res, err = s.service.ListOrders(ctx, r.Skip, r.Take)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %w", err)
	}

	orders := []*pb.Order{}
	for _, order := range res {
		orders = append(orders, &pb.Order{
			Id:        order.ID,
			AccountID: order.AccountID,
			CreatedAt: order.CreatedAt.Format(time.RFC3339),
			UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
			Phase:     order.Phase,
		})
	}

	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}

// UpdateOrder handles the gRPC call to update an order
func (s *grpcServer) UpdateOrder(ctx context.Context, r *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	order := &Order{
		ID:        r.Order.Id,
		AccountID: r.Order.AccountID,
		Phase:     r.Order.Phase,
		UpdatedAt: time.Now(),
	}

	updatedOrder, err := s.service.UpdateOrder(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	return &pb.UpdateOrderResponse{
		Order: &pb.Order{
			Id:        updatedOrder.ID,
			AccountID: updatedOrder.AccountID,
			CreatedAt: updatedOrder.CreatedAt.Format(time.RFC3339),
			UpdatedAt: updatedOrder.UpdatedAt.Format(time.RFC3339),
			Phase:     updatedOrder.Phase,
		},
	}, nil
}
