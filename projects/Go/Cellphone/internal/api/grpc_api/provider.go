package grpc_api

import (
	pb "cellphone/protos/go"
	"context"
	"errors"
)

func (self *GRPCServer) GetProviderById(ctx context.Context, req *pb.GetProviderByIdRequest) (*pb.GetProviderByIdResponse, error) {
	id := req.GetId()

	entity, err := self.repo.Provider.GetById(int(id))

	if err != nil {
		return nil, err
	}

	return &pb.GetProviderByIdResponse{
		Provider: &pb.Provider{
			Id:    entity.Id,
			Name:  entity.Name,
			Total: entity.Total,
		},
	}, nil
}

func (self *GRPCServer) GetProviderByName(ctx context.Context, req *pb.GetProviderByNameRequest) (*pb.GetProviderByNameResponse, error) {
	return nil, errors.New("Unimplemented")
}

func (self *GRPCServer) InsertProvider(ctx context.Context, req *pb.InsertProviderRequest) (*pb.InsertProviderResponse, error) {
	return nil, errors.New("Unimplemented")
}

func (self *GRPCServer) DeleteProvider(ctx context.Context, req *pb.DeleteProviderRequest) (*pb.DeleteProviderResponse, error) {
	return nil, errors.New("Unimplemented")
}

func (self *GRPCServer) UpdateProvider(ctx context.Context, req *pb.UpdateProviderRequest) (*pb.UpdateProviderResponse, error) {
	return nil, errors.New("Unimplemented")
}
