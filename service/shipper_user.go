package service

import (
	"context"
	"fmt"
	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

// ShipperUserService ...
type ShipperUserService struct {
	storage storage.StorageI
	logger  l.Logger
}

// NewShipperUserService ...
func NewShipperUserService(strg storage.StorageI, log l.Logger) *ShipperUserService {
	return &ShipperUserService{
		storage: strg,
		logger:  log,
	}
}

// Create is function for creating a courier
func (s *ShipperUserService) Create(ctx context.Context, req *pb.ShipperUser) (*pb.ShipperUserId, error) {
	shipperUserID, err := s.storage.ShipperUser().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating shipper user", req)
	}

	return &pb.ShipperUserId{
		Id: shipperUserID,
	}, nil
}

// Get is function for getting a ShipperUser
func (s *ShipperUserService) Get(ctx context.Context, req *pb.ShipperUserId) (*pb.ShipperUser, error) {
	shipperUser, err := s.storage.ShipperUser().Get(req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting shipper user", req)
	}

	return shipperUser, nil
}

// GetAll is function for getting all ShipperUsers
func (s *ShipperUserService) GetAll(ctx context.Context, req *pb.GetAllShipperUsersRequest) (*pb.GetAllShipperUsersResponse, error) {
	shipperUsers, count, err := s.storage.ShipperUser().GetAll(req.GetPage(), req.GetLimit(), req.ShipperId, req.GetUserRoleId(), req.GetSearch())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all shipper users", req)
	}

	return &pb.GetAllShipperUsersResponse{
		ShipperUsers: shipperUsers,
		Count:        count,
	}, nil
}

// Update is function for updating a ShipperUser
func (s *ShipperUserService) Update(ctx context.Context, req *pb.ShipperUser) (*gpb.Empty, error) {
	err := s.storage.ShipperUser().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating shipper user", req)
	}

	return &gpb.Empty{}, nil
}

//Delete if function for deleting ShipperUser
func (s *ShipperUserService) Delete(ctx context.Context, req *pb.DeleteShipperUserRequest) (*gpb.Empty, error) {
	err := s.storage.ShipperUser().Delete(req.Id, req.ShipperId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting shipper user", req)
	}

	return &gpb.Empty{}, nil
}

// GetByUsername is function for getting shipperUser by loging(username)
func (s *ShipperUserService) GetByUsername(ctx context.Context, req *pb.GetShipperUserByUsernameRequest) (*pb.ShipperUser, error) {
	shipperUser, err := s.storage.ShipperUser().GetByUsername(req.Username)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting shipper user by username", req)
	}

	return shipperUser, nil
}

// GetByCredentials is function for getting shipperUser by credentials
func (s *ShipperUserService) GetByCredentials(ctx context.Context, req *pb.GetByCredentialsRequest) (*pb.ShipperUser, error) {
	shipperUser, err := s.storage.ShipperUser().GetByUsername(req.Username)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting shipper user by login", req)
	}

	err = bcrypt.CompareHashAndPassword([]byte(shipperUser.Password), []byte(req.Password))
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s, Not found", err), l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	}

	return shipperUser, nil
}

// ChangePassword is function for changing password
func (s *ShipperUserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*gpb.Empty, error) {
	err := s.storage.ShipperUser().ChangePassword(req.Id, req.Password)
	if err != nil {
		return nil, handleError(s.logger, err, "error while changing shipper user password", req)
	}

	return &gpb.Empty{}, nil
}
