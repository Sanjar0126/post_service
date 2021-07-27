package repo

import (
	pb "genproto/user_service"
)

//BranchUserStorageI ...
type BranchUserStorageI interface {
	Create(branchUser *pb.BranchUser) (string, error)
	Update(branchUser *pb.BranchUser) error
	Delete(id, shipperId string) error
	Get(id string) (*pb.BranchUser, error)
	GetAll(page, limit uint64, shipperId, search, branchId, userRoleID string) ([]*pb.BranchUser, uint64, error)
	UpdateFcmToken(id, shipperId, fcmToken, platformID string) error
	DeleteFcmToken(id string) error
	GetByPhone(phone, shipperId string) (*pb.BranchUser, error)
}
