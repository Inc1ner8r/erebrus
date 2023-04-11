package status

import (
	"context"
	"errors"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/model"
	log "github.com/sirupsen/logrus"
)

type StatusService struct {
	UnimplementedStatusServiceServer
}

func (s *StatusService) GetStatusGrpc(ctx context.Context, request *Empty) (*model.Status, error) {
	status_data, err := core.GetServerStatus()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Failed to get server status")
		return nil, errors.New(err.Error())
	}

	statusProto := &model.Status{
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
		PersistentKeepalive: status_data.PersistentKeepalive,
		DNS:                 status_data.DNS,
	}

	//test fields
	// protobyte, _ := proto.Marshal(statusProto)
	// fmt.Println(protobyte)
	// newMessage := &model.Status{}
	// fmt.Println(proto.Unmarshal(protobyte, newMessage))
	// fmt.Println(newMessage)

	return statusProto, nil
}

//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative gRPC/v1/status/status.proto
