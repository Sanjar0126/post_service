package service

import (
	"context"
	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/crypto/bcrypt"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

// SystemUserService ...
type SystemUserService struct {
	storage storage.StorageI
	logger  l.Logger
}

// NewSystemUserService ...
func NewSystemUserService(strg storage.StorageI, log l.Logger) *SystemUserService {
	return &SystemUserService{
		storage: strg,
		logger:  log,
	}
}

// Create is function for creating a courier
func (s *SystemUserService) Create(ctx context.Context, req *pb.SystemUser) (*pb.SystemUserId, error) {
	systemUserID, err := s.storage.SystemUser().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating system user", req)
	}

	return &pb.SystemUserId{
		Id: systemUserID,
	}, nil
}

// Get is function for getting a systemUser
func (s *SystemUserService) Get(ctx context.Context, req *pb.SystemUserId) (*pb.SystemUser, error) {
	systemUser, err := s.storage.SystemUser().Get(req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting system user", req)
	}

	return systemUser, nil
}

// GetAll is function for getting all systemUsers
func (s *SystemUserService) GetAll(ctx context.Context, req *pb.GetAllSystemUsersRequest) (*pb.GetAllSystemUsersResponse, error) {
	systemUsers, count, err := s.storage.SystemUser().GetAll(req.GetPage(), req.GetLimit(), req.GetSearch())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all system users", req)
	}

	return &pb.GetAllSystemUsersResponse{
		SystemUsers: systemUsers,
		Count:       count,
	}, nil
}

// Update is function for updating a systemUser
func (s *SystemUserService) Update(ctx context.Context, req *pb.SystemUser) (*gpb.Empty, error) {
	err := s.storage.SystemUser().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating system user", req)
	}

	return &gpb.Empty{}, nil
}

//Delete if function for deleting systemUser
func (s *SystemUserService) Delete(ctx context.Context, req *pb.DeleteSystemUserRequest) (*gpb.Empty, error) {
	err := s.storage.SystemUser().Delete(req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting system user", req)
	}

	return &gpb.Empty{}, nil
}

// GetByUsername is a function to get a SystemUser by login
func (s *SystemUserService) GetByUsername(ctx context.Context, req *pb.GetSystemUserByUsernameRequest) (*pb.SystemUser, error) {
	systemUser, err := s.storage.SystemUser().GetByUsername(req.Username)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting system user by username", req)
	}

	return systemUser, nil
}

//GetByCredentials is a function to get systemUser by credentials
func (s *SystemUserService) GetByCredentials(ctx context.Context, req *pb.GetSystemUserByCredentialsRequest) (*pb.SystemUser, error) {
	systemUser, err := s.storage.SystemUser().GetByUsername(req.Username)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting systemUser by login", req)
	}

	err = bcrypt.CompareHashAndPassword([]byte(systemUser.Password), []byte(req.Password))
	if err != nil {
		return nil, handleError(s.logger, err, "error while comparing hash and password of system user", req)
	}

	return systemUser, nil
}

// ChangePassword is a function to change systemUsers password
func (s *SystemUserService) ChangePassword(ctx context.Context, req *pb.ChangeSystemUserPasswordRequest) (*gpb.Empty, error) {
	err := s.storage.SystemUser().ChangePassword(req.Id, req.Password)
	if err != nil {
		return nil, handleError(s.logger, err, "error while changing system user password", req)
	}

	return &gpb.Empty{}, nil
}
