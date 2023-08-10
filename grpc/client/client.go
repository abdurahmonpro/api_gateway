package client

import (
	"api_gateway/config"
	"api_gateway/genproto/auth_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	UserService() auth_service.UserServiceClient
	AuthService() auth_service.AuthServiceClient
}

type grpcClients struct {
	authService auth_service.AuthServiceClient
	userServie auth_service.UserServiceClient
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

	return &grpcClients{
		authService: auth_service.NewAuthServiceClient(connAuthService),
		userServie: auth_service.NewUserServiceClient(connUserService),
	}, nil
}

func (g *grpcClients) UserService() auth_service.UserServiceClient {
	return g.userServie
}

func (g *grpcClients) AuthService() auth_service.AuthServiceClient {
	return g.authService
}