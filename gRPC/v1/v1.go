package v1

import (
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/client"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/paseto"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/server"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/status"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/test"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
)

func Initialize() *grpc.Server {

	//get the instance of server and client services
	ServerService := &server.ServerService{}
	ClientService := &client.ClientService{}
	StatusService := &status.StatusService{}
	TestService := &test.TestService{}

	//creating a new gRPC server
	grpc_server := grpc.NewServer(
		grpc.StreamInterceptor(auth.StreamServerInterceptor(paseto.PASETO)),
		grpc.UnaryInterceptor(auth.UnaryServerInterceptor(paseto.PASETO)),
	)

	server.RegisterServerServiceServer(grpc_server, ServerService)
	client.RegisterClientServiceServer(grpc_server, ClientService)
	status.RegisterStatusServiceServer(grpc_server, StatusService)
	test.RegisterTestServiceServer(grpc_server, TestService)

	return grpc_server
}
