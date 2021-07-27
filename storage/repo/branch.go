package repo

import (
	pb "genproto/user_service"

	_ "github.com/lib/pq"
)

//BranchStorageI ...
type BranchStorageI interface {
	Create(branch *pb.Branch) (string, error)
	Update(branch *pb.Branch) error
	Delete(id, shipperId string) error
	Get(id string) (*pb.Branch, error)
	GetAll(shipperId, search, fareId string, page, limit uint64, jowi, iiko bool) ([]*pb.Branch, uint64, error)
	GetNearestBranch(shipperId string, location *pb.Location) ([]*pb.Branch, error)
	GetByName(shipperId, name string) (*pb.Branch, error)
	GetByJowiID(jowiID string) (*pb.Branch, error)
	GetByIikoID(iikoID string) (*pb.Branch, error)
}
