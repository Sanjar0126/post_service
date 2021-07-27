package service

import (
	"context"

	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

type IikoCredentialsService struct {
	storage storage.StorageI
	logger  l.Logger
}

func NewIikoCredentialsService(strg storage.StorageI, log l.Logger) *IikoCredentialsService {
	return &IikoCredentialsService{
		storage: strg,
		logger:  log,
	}
}

func (is *IikoCredentialsService) Create(ctx context.Context, req *pb.IikoCredentials) (*pb.IikoCredentialsId, error) {
	shipperID, err := is.storage.IikoCredentials().Create(req)

	if err != nil {
		return nil, handleError(is.logger, err, "error while creating iiko_credentials", req)
	}

	return &pb.IikoCredentialsId{ShipperId: shipperID}, nil
}

func (is *IikoCredentialsService) Update(ctx context.Context, req *pb.IikoCredentials) (*gpb.Empty, error) {
	err := is.storage.IikoCredentials().Update(req)

	if err != nil {
		return nil, handleError(is.logger, err, "error while updating iiko_credentials", req)
	}

	return &gpb.Empty{}, nil
}

func (is *IikoCredentialsService) Delete(ctx context.Context, req *pb.IikoCredentialsId) (*gpb.Empty, error) {
	err := is.storage.IikoCredentials().Delete(req.GetShipperId())

	if err != nil {
		return nil, handleError(is.logger, err, "error while deleting iiko_credentials", req)
	}

	return &gpb.Empty{}, nil
}

func (is *IikoCredentialsService) Get(ctx context.Context, req *pb.IikoCredentialsId) (*pb.IikoCredentials, error) {
	iikoC, err := is.storage.IikoCredentials().Get(req.GetShipperId())

	if err != nil {
		return nil, handleError(is.logger, err, "error while getting iiko_credentials", req)
	}

	return iikoC, nil
}
