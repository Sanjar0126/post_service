package repo

import (
	pb "genproto/user_service"
)

//SystemUserStorageI ...
type SystemUserStorageI interface {
	Create(systemUser *pb.SystemUser) (string, error)
	Update(systemUser *pb.SystemUser) error
	Delete(id string) error
	Get(id string) (*pb.SystemUser, error)
	GetAll(page, limit uint64, search string) ([]*pb.SystemUser, uint64, error)
	GetByUsername(username string) (*pb.SystemUser, error)
	ChangePassword(id, password string) error
}
