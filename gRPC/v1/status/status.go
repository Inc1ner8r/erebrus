package status

import (
	"context"
	"fmt"

	"github.com/TheLazarusNetwork/erebrus/core"
	"google.golang.org/protobuf/proto"
)

type StatusService struct {
	UnimplementedStatusServiceServer
}

func (s *StatusService) GetStatusGrpc(ctx context.Context, request *Empty) (*Status, error) {
	status_data, err := core.GetServerStatus()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	statusProto := &Status{
		Version:             status_data.Version,
		Hostname:            status_data.Hostname,
		Domain:              status_data.Domain,
		PublicIP:            status_data.PublicIP,
		GRPCPort:            status_data.GRPCPort,
		PrivateIP:           status_data.PrivateIP,
		HttpPort:            status_data.HttpPort,
		Region:              status_data.Region,
		VPNPort:             status_data.VPNPort,
		PublicKey:           status_data.PublicKey,
		PersistentKeepalive: uint32(status_data.PersistentKeepalive),
		DNS:                 status_data.DNS,
	}
	protobyte, _ := proto.Marshal(statusProto)
	fmt.Println(protobyte)
	newMessage := &Status{}
	fmt.Println(proto.Unmarshal(protobyte, newMessage))
	fmt.Println(newMessage)
	return statusProto, nil

}

//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative gRPC/v1/status/status.proto
