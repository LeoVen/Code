package grpc_api

import (
	"context"
	"errors"

	pb "cellphone/protos/go"
)

func (self *GRPCServer) GetCellphoneById(ctx context.Context, req *pb.IdRequest) (*pb.Cellphone, error) {
	id := req.GetId()

	entity, err := self.repo.Cellphone.GetById(int(id))

	if err != nil {
		return nil, err
	}

	return &pb.Cellphone{
		Id:         entity.Id,
		ProviderId: entity.ProviderId,
		Number:     entity.Number,
	}, nil
}

func (self *GRPCServer) FetchSingle(ctx context.Context, req *pb.FetchSingleRequest) (*pb.FetchSingleResponse, error) {
	return nil, errors.New("Unimplemented")
}

func (self *GRPCServer) BulkInsert(ctx context.Context, req *pb.BulkInsertRequest) (*pb.BulkInsertResponse, error) {
	return nil, errors.New("Unimplemented")
}
