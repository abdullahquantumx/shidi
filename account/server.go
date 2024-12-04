package account

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/Shridhar2104/logilo/account/pb"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

func NewGRPCServer(service Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &grpcServer{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{}, // Add this
		service: service,
	})
	reflection.Register(server)
	return server.Serve(lis)
}

func (s *grpcServer) CreateAccount(ctx context.Context, r *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	a, err := s.service.CreateAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{Account: &pb.Account{
		Id: a.ID,
		Name: a.Name,
	}}, nil
}


func (s *grpcServer) GetAccountByID(ctx context.Context, r *pb.GetAccountByIDRequest) (*pb.GetAccountByIDResponse, error) {
	a, err := s.service.GetAccountByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountByIDResponse{Account: &pb.Account{
		Id: a.ID,
		Name: a.Name,
	}}, nil	
}


func (s *grpcServer) ListAccounts(ctx context.Context, r *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	accounts, err := s.service.ListAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	a:= []*pb.Account{}
	for _, account := range accounts {
		a = append(a, &pb.Account{
			Id: account.ID,
			Name: account.Name,
		})
	} 
	return &pb.ListAccountsResponse{Accounts: a}, nil
}
