package service

import (
	"context"

	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

type JowiCredentialsService struct {
	storage storage.StorageI
	logger  l.Logger
}

func NewJowiCredentialsService(strg storage.StorageI, log l.Logger) *JowiCredentialsService {
	return &JowiCredentialsService{
		storage: strg,
		logger:  log,
	}
}

func (js *JowiCredentialsService) Create(ctx context.Context, req *pb.JowiCredentials) (*pb.JowiCredentialsId, error) {
	shipperID, err := js.storage.JowiCredentials().Create(req)

	if err != nil {
		return nil, handleError(js.logger, err, "error while creating jowi_credentials", req)
	}

	return &pb.JowiCredentialsId{ShipperId: shipperID}, nil
}

func (js *JowiCredentialsService) Update(ctx context.Context, req *pb.JowiCredentials) (*gpb.Empty, error) {
	err := js.storage.JowiCredentials().Update(req)

	if err != nil {
		return nil, handleError(js.logger, err, "error while updating jowi_credentials", req)
	}

	return &gpb.Empty{}, nil
}

func (js *JowiCredentialsService) Delete(ctx context.Context, req *pb.JowiCredentialsId) (*gpb.Empty, error) {
	err := js.storage.JowiCredentials().Delete(req.GetShipperId())

	if err != nil {
		return nil, handleError(js.logger, err, "error while deleting jowi_credentials", req)
	}

	return &gpb.Empty{}, nil
}

func (js *JowiCredentialsService) Get(ctx context.Context, req *pb.JowiCredentialsId) (*pb.JowiCredentials, error) {
	jowiC, err := js.storage.JowiCredentials().Get(req.GetShipperId())

	if err != nil {
		return nil, handleError(js.logger, err, "error while getting jowi_credentials", req)
	}

	return jowiC, nil
}
