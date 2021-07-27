package repo

import (
	pb "genproto/user_service"
)

type CustomerTypeI interface {
	Create(customer_type *pb.CustomerType) (string, error)
	Update(customer_type *pb.CustomerType) error
	Get(id string) (*pb.CustomerType, error)
	GetAll(page, limit int64, name string) ([]*pb.CustomerType, int64, error)
	Delete(id string) error
}
