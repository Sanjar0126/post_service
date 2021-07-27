package repo

import (
	pb "genproto/user_service"
)

//ShipperStorageI ...
type ShipperStorageI interface {
	Create(shipper *pb.Shipper) (string, error)
	Update(shipper *pb.Shipper) error
	Delete(id string) error
	Get(id string) (*pb.Shipper, error)
	GetAll(page, limit uint64, hasIiko bool) ([]*pb.Shipper, uint64, error)
	GetByName(name string) (*pb.Shipper, error)
}
