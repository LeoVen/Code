package grpc_api

import (
	pb "cellphone/protos/go"
	"context"
)

func (self *GRPCServer) GetProviderById(ctx context.Context, req *pb.IdRequest) (*pb.Provider, error) {
	id := req.GetId()

	entity, err := self.repo.Provider.GetById(int(id))

	if err != nil {
		return nil, err
	}

	return &pb.Provider{
		Id:    entity.Id,
		Name:  entity.Name,
		Total: entity.Total,
	}, nil
}
