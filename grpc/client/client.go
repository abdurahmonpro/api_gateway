package client

import (
	"api_gateway/config"
	"api_gateway/genproto/auth_service"
	"api_gateway/genproto/book_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	UserService() auth_service.UserServiceClient
	AuthService() auth_service.AuthServiceClient
	BookService() book_service.BookServiceClient
}

type grpcClients struct {
	authService auth_service.AuthServiceClient
	userServie  auth_service.UserServiceClient
	bookService book_service.BookServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connAuthService, err := grpc.Dial(
		cfg.AuthServiceHost+cfg.AuthGRPCPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	connUserService, err := grpc.Dial(
		cfg.AuthServiceHost+cfg.AuthGRPCPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	connBookService, err := grpc.Dial(
		cfg.BookServiceHost+cfg.BookGRPCPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		authService: auth_service.NewAuthServiceClient(connAuthService),
		userServie:  auth_service.NewUserServiceClient(connUserService),
		bookService: book_service.NewBookServiceClient(connBookService),
	}, nil
}

func (g *grpcClients) UserService() auth_service.UserServiceClient {
	return g.userServie
}

func (g *grpcClients) AuthService() auth_service.AuthServiceClient {
	return g.authService
}

func (g *grpcClients) BookService() book_service.BookServiceClient {
	return g.bookService
}
