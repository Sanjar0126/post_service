package repo

import (
	pb "genproto/user_service"
)

type ClickInfoStorageI interface {
	Create(click *pb.Click) (string, error)
	Get(id, branchID string) (*pb.Click, error)
	GetAll(shipperID string, branchIDs []string, page, limit uint64) ([]*pb.Click, uint64, error)
	Update(click *pb.Click) error
	Delete(id, branchID string) error
	GetShipperAndKeyByCredentials(serviceID int64) (string, string, error)
}
