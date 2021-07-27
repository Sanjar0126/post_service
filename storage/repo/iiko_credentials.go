package repo

import (
	pb "genproto/user_service"

	_ "github.com/lib/pq"
)

//JowiCredentialsStorageI ...
type IikoCredentialsStorageI interface {
	Create(iikoC *pb.IikoCredentials) (string, error)
	Update(iikoC *pb.IikoCredentials) error
	Delete(shipperID string) error
	Get(shipperID string) (*pb.IikoCredentials, error)
}
