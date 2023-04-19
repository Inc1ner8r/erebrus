package v1

import (
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/authenticate"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/authenticate/challengeid"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/authenticate/paseto"
	selector_grpc "github.com/TheLazarusNetwork/erebrus/gRPC/v1/authenticate/selector"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/client"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/server"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/status"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"google.golang.org/grpc"
)

func Initialize() *grpc.Server {

	//get the instance of server and client services
	ServerService := &server.ServerService{}
	ClientService := &client.ClientService{}
	StatusService := &status.StatusService{}
	AuthenticateService := &authenticate.AuthenticateService{}
	ChallengeIdService := &challengeid.ChallengeIdService{}

	//creating a new gRPC server
	grpc_server := grpc.NewServer(
		grpc.ChainStreamInterceptor(selector.StreamServerInterceptor(auth.StreamServerInterceptor(paseto.PASETO), selector.MatchFunc(selector_grpc.LoginSkip))),
		grpc.ChainUnaryInterceptor(selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(paseto.PASETO), selector.MatchFunc(selector_grpc.LoginSkip))),
	)
	server.RegisterServerServiceServer(grpc_server, ServerService)
	client.RegisterClientServiceServer(grpc_server, ClientService)
	status.RegisterStatusServiceServer(grpc_server, StatusService)
	authenticate.RegisterAuthenticateServiceServer(grpc_server, AuthenticateService)
	challengeid.RegisterChallengeidServiceServer(grpc_server, ChallengeIdService)

	return grpc_server
}
