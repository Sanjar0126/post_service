package service

import (
	"context"
	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

type PaymeInfoService struct {
	storage storage.StorageI
	logger  l.Logger
}

func NewPaymeInfoService(strg storage.StorageI, log l.Logger) *PaymeInfoService {
	return &PaymeInfoService{
		storage: strg,
		logger:  log,
	}
}

func (s *PaymeInfoService) Create(ctx context.Context, req *pb.Payme) (*pb.PaymeInfoId, error) {
	shipperID, err := s.storage.PaymeInfo().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating payme info", req)
	}

	return &pb.PaymeInfoId{
		Id: shipperID,
	}, nil
}

func (s *PaymeInfoService) Get(ctx context.Context, req *pb.PaymeInfoId) (*pb.Payme, error) {
	paymeInfo, err := s.storage.PaymeInfo().Get(req.Id, req.BranchId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting payme info", req)
	}

	if paymeInfo.BranchId != "" {
		branch, err := s.storage.Branch().Get(paymeInfo.BranchId)
		if err != nil {
			return nil, handleError(s.logger, err, "error while getting branch name", req)
		}

		paymeInfo.BranchName = branch.Name
	}

	return paymeInfo, nil
}

func (s *PaymeInfoService) GetAll(ctx context.Context, req *pb.GetAllPaymeInfoRequest) (*pb.GetAllPaymeInfoResponse, error) {
	var branchIDs []string
	branchesList := make(map[string]string)

	branches, _, err := s.storage.Branch().GetAll(req.ShipperId, req.Search, "", 1, 100, false, false)
	if err != nil {
		return nil, handleError(s.logger, err, "Error while getting all branches", req)
	}

	for _, branch := range branches {
		branchIDs = append(branchIDs, branch.Id)
		branchesList[branch.Id] = branch.Name
	}

	paymeInfos, count, err := s.storage.PaymeInfo().GetAll(req.ShipperId, branchIDs, req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting payme info", req)
	}

	for _, paymeInfo := range paymeInfos {
		if paymeInfo.BranchId != "" {
			paymeInfo.BranchName = branchesList[paymeInfo.BranchId]
		}
	}

	return &pb.GetAllPaymeInfoResponse{
		PaymeInfos: paymeInfos,
		Count:      count,
	}, nil
}

func (s *PaymeInfoService) Update(ctx context.Context, req *pb.Payme) (*gpb.Empty, error) {
	err := s.storage.PaymeInfo().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while upadting payme info", req)
	}

	return &gpb.Empty{}, nil
}

func (s *PaymeInfoService) Delete(ctx context.Context, req *pb.PaymeInfoId) (*gpb.Empty, error) {
	err := s.storage.PaymeInfo().Delete(req.Id, req.BranchId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting payme info", req)
	}

	return &gpb.Empty{}, nil
}

func (s *PaymeInfoService) GetShipperByCredentials(ctx context.Context, req *pb.PaymeCredentials) (*pb.PaymeInfoId, error) {
	shipperID, err := s.storage.PaymeInfo().GetShipperByCredentials(req.Token)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting shipper by payme token", req)
	}

	return &pb.PaymeInfoId{
		Id: shipperID,
	}, nil
}
