package repo

import (
	pb "genproto/user_service"

	_ "github.com/lib/pq"
)

//JowiCredentialsStorageI ...
type JowiCredentialsStorageI interface {
	Create(jowiC *pb.JowiCredentials) (string, error)
	Update(jowiC *pb.JowiCredentials) error
	Delete(shipperID string) error
	Get(shipperID string) (*pb.JowiCredentials, error)
}
