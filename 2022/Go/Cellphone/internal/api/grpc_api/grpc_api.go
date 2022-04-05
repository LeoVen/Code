package grpc_api

import (
	"cellphone/internal/app_config"
	"cellphone/internal/repository"
	"cellphone/internal/telemetry"
	"net"

	"google.golang.org/grpc"

	pb "cellphone/protos/go"
)

type GRPCServer struct {
	pb.UnimplementedServiceServer
	repo *repository.RepositoryService
	tel  *telemetry.Telemetry
}

func (self *GRPCServer) Start(config app_config.Main) error {
	listener, err := net.Listen("tcp", ":"+config.Flags["CELL_APIPORT"])

	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	pb.RegisterServiceServer(srv, self)

	return srv.Serve(listener)
}

func NewServer(repo *repository.RepositoryService, tel *telemetry.Telemetry) *GRPCServer {
	return &GRPCServer{
		repo: repo,
		tel:  tel,
	}
}
