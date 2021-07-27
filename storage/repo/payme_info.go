package repo

import (
	pb "genproto/user_service"
)

type PaymeInfoStorageI interface {
	Create(payme *pb.Payme) (string, error)
	Get(id string, branchID string) (*pb.Payme, error)
	GetAll(shipperID string, branchIDs []string, page, limit uint64) ([]*pb.Payme, uint64, error)
	Update(payme *pb.Payme) error
	Delete(id, branchID string) error
	GetShipperByCredentials(token string) (string, error)
}
