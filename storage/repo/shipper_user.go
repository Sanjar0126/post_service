package repo

import (
	pb "genproto/user_service"
)

//ShipperUserStorageI ...
type ShipperUserStorageI interface {
	Create(shipperUser *pb.ShipperUser) (string, error)
	Update(shipperUser *pb.ShipperUser) error
	Delete(id, shipperId string) error
	Get(id string) (*pb.ShipperUser, error)
	GetAll(page, limit uint64, shipperId, userRoleID, search string) ([]*pb.ShipperUser, uint64, error)
	GetByUsername(username string) (*pb.ShipperUser, error)
	ChangePassword(id, password string) error
}
