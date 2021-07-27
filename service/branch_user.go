package service

import (
	"context"
	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

// BranchUserService ...
type BranchUserService struct {
	storage storage.StorageI
	logger  l.Logger
}

// NewBranchUserService ...
func NewBranchUserService(strg storage.StorageI, log l.Logger) *BranchUserService {
	return &BranchUserService{
		storage: strg,
		logger:  log,
	}
}

// Create is function for creating a BranchUser
func (s *BranchUserService) Create(ctx context.Context, req *pb.BranchUser) (*pb.BranchUserId, error) {
	branchUserID, err := s.storage.BranchUser().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating branch user", req)
	}

	return &pb.BranchUserId{
		Id: branchUserID,
	}, nil
}

// Get is function for getting a BranchUser
func (s *BranchUserService) Get(ctx context.Context, req *pb.BranchUserId) (*pb.BranchUser, error) {
	branchUser, err := s.storage.BranchUser().Get(req.GetId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting branch user", req)
	}

	return branchUser, nil
}

// GetAll is function for getting all BranchUsers
func (s *BranchUserService) GetAll(ctx context.Context, req *pb.GetAllBranchUsersRequest) (*pb.GetAllBranchUsersResponse, error) {
	branchUsers, count, err := s.storage.BranchUser().GetAll(req.GetPage(), req.GetLimit(), req.GetShipperId(), req.GetSearch(), req.GetBranchId(), req.GetUserRoleId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all branch users", req)
	}

	return &pb.GetAllBranchUsersResponse{
		BranchUsers: branchUsers,
		Count:       count,
	}, nil
}

// Update is function for updating a BranchUser
func (s *BranchUserService) Update(ctx context.Context, req *pb.BranchUser) (*gpb.Empty, error) {
	err := s.storage.BranchUser().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating branch user", req)
	}

	return &gpb.Empty{}, nil
}

//Delete is function for deleting BranchUser
func (s *BranchUserService) Delete(ctx context.Context, req *pb.DeleteBranchUserRequest) (*gpb.Empty, error) {
	err := s.storage.BranchUser().Delete(req.Id, req.ShipperId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting branch user", req)
	}

	return &gpb.Empty{}, nil
}

// UpdateFcmToken is function to update Fcm token
func (s *BranchUserService) UpdateFcmToken(ctx context.Context, req *pb.UpdateBranchUserFcmTokenRequest) (*gpb.Empty, error) {
	err := s.storage.BranchUser().UpdateFcmToken(req.Id, req.ShipperId, req.FcmToken, req.PlatformId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating branch user fcmToken", req)
	}

	return &gpb.Empty{}, nil
}

func (s *BranchUserService) DeleteFcmToken(ctx context.Context, req *pb.BranchUserId) (*gpb.Empty, error) {
	err := s.storage.BranchUser().DeleteFcmToken(req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting branch user fcm token", req)
	}

	return &gpb.Empty{}, nil
}

func (s *BranchUserService) GetByPhone(ctx context.Context, req *pb.GetBranchUserByPhoneRequest) (*pb.BranchUser, error) {
	branchUser, err := s.storage.BranchUser().GetByPhone(req.Phone, req.ShipperId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting branch user by phone", req)
	}

	return branchUser, nil
}
