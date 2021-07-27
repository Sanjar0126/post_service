package service

import (
	"context"
	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

type ClickInfoService struct {
	storage storage.StorageI
	logger  l.Logger
}

func NewClickInfoService(strg storage.StorageI, log l.Logger) *ClickInfoService {
	return &ClickInfoService{
		storage: strg,
		logger:  log,
	}
}

func (s *ClickInfoService) Create(ctx context.Context, req *pb.Click) (*pb.ClickInfoId, error) {
	shipperID, err := s.storage.ClickInfo().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating click info", req)
	}

	return &pb.ClickInfoId{
		Id: shipperID,
	}, nil
}

func (s *ClickInfoService) Get(ctx context.Context, req *pb.ClickInfoId) (*pb.Click, error) {
	clickInfo, err := s.storage.ClickInfo().Get(req.Id, req.BranchId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting click info", req)
	}

	if clickInfo.BranchId != "" {
		branch, err := s.storage.Branch().Get(clickInfo.BranchId)
		if err != nil {
			return nil, handleError(s.logger, err, "error while getting branch name", req)
		}

		clickInfo.BranchName = branch.Name
	}
	return clickInfo, nil
}

func (s *ClickInfoService) GetAll(ctx context.Context, req *pb.GetAllClickInfoRequest) (*pb.GetAllClickInfoResponse, error) {
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

	clickInfos, count, err := s.storage.ClickInfo().GetAll(req.ShipperId, branchIDs, req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting click info", req)
	}

	for _, clickInfo := range clickInfos {
		if clickInfo.BranchId != "" {
			clickInfo.BranchName = branchesList[clickInfo.BranchId]
		}
	}

	return &pb.GetAllClickInfoResponse{
		ClickInfos: clickInfos,
		Count:      count,
	}, nil
}

func (s *ClickInfoService) Update(ctx context.Context, req *pb.Click) (*gpb.Empty, error) {
	err := s.storage.ClickInfo().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while upadting payme info", req)
	}

	return &gpb.Empty{}, nil
}

func (s *ClickInfoService) Delete(ctx context.Context, req *pb.ClickInfoId) (*gpb.Empty, error) {
	err := s.storage.ClickInfo().Delete(req.Id, req.BranchId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting payme info", req)
	}

	return &gpb.Empty{}, nil
}

func (s *ClickInfoService) GetShipperAndKeyByCredentials(ctx context.Context, req *pb.ClickCredentials) (*pb.ClickShipperIdAndKey, error) {
	shipperID, key, err := s.storage.ClickInfo().GetShipperAndKeyByCredentials(req.ServiceId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting shipper and key by click credentials", req)
	}

	return &pb.ClickShipperIdAndKey{
		Id:  shipperID,
		Key: key,
	}, nil
}
