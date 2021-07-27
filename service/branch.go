package service

import (
	"context"
	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

// BranchService ...
type BranchService struct {
	storage storage.StorageI
	logger  l.Logger
}

// NewBranchService ...
func NewBranchService(strg storage.StorageI, log l.Logger) *BranchService {
	return &BranchService{
		storage: strg,
		logger:  log,
	}
}

// Create is function for creating a courier
func (s *BranchService) Create(ctx context.Context, req *pb.Branch) (*pb.BranchId, error) {
	branchId, err := s.storage.Branch().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating branch", req)
	}

	return &pb.BranchId{
		Id: branchId,
	}, nil
}

// Get is function for getting a branch
func (s *BranchService) Get(ctx context.Context, req *pb.BranchId) (*pb.Branch, error) {
	branch, err := s.storage.Branch().Get(req.GetId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting branch", req)
	}

	return branch, nil
}

// GetAll is function for getting all couriers
func (s *BranchService) GetAll(ctx context.Context, req *pb.GetAllBranchesRequest) (*pb.GetAllBranchesResponse, error) {
	branches, count, err := s.storage.Branch().GetAll(req.ShipperId, req.GetSearch(), req.GetFareId(), req.GetPage(), req.GetLimit(), req.GetJowi(), req.GetIiko())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all branch", req)
	}

	return &pb.GetAllBranchesResponse{
		Branches: branches,
		Count:    count,
	}, nil
}

// Update is function for updating a branch
func (s *BranchService) Update(ctx context.Context, req *pb.Branch) (*gpb.Empty, error) {
	err := s.storage.Branch().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating branch", req)
	}

	return &gpb.Empty{}, nil
}

//Delete if function for deleting branch
func (s *BranchService) Delete(ctx context.Context, req *pb.DeleteBranchRequest) (*gpb.Empty, error) {
	err := s.storage.Branch().Delete(req.GetId(), req.GetShipperId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting branch", req)
	}

	return &gpb.Empty{}, nil
}

func (s *BranchService) GetNearestBranch(ctx context.Context, req *pb.GetNearestBranchRequest) (*pb.GetNearestBranchResponse, error) {
	branches, err := s.storage.Branch().GetNearestBranch(req.ShipperId, req.GetLocation())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting nearest branch", req)
	}

	return &pb.GetNearestBranchResponse{
		Branches: branches,
	}, nil
}

func (s *BranchService) GetByName(ctx context.Context, req *pb.GetBranchByNameRequest) (*pb.Branch, error) {
	branch, err := s.storage.Branch().GetByName(req.ShipperId, req.Name)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting branch by name", req)
	}

	return branch, nil
}

func (s *BranchService) GetByJowiID(ctx context.Context, req *pb.GetByJowiIDRequest) (*pb.Branch, error) {
	branch, err := s.storage.Branch().GetByJowiID(req.GetJowiId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting branch by jowi id", req)
	}

	return branch, nil
}

func (s *BranchService) GetByIikoID(ctx context.Context, req *pb.GetByIikoIDRequest) (*pb.Branch, error) {
	branch, err := s.storage.Branch().GetByIikoID(req.GetIikoId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting branch by iiko id", req)
	}

	return branch, nil
}
