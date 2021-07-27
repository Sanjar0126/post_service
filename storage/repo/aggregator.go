package repo

import (
	pb "genproto/user_service"

	_ "github.com/lib/pq"
)

//AggregatorStorageI ...
type AggregatorStorageI interface {
	Create(aggregator *pb.Aggregator) (string, error)
	Update(aggregator *pb.Aggregator) error
	Delete(id string) error
	Get(id string) (*pb.Aggregator, error)
	GetAll(page, limit int64, name string, shipper_id string) ([]*pb.Aggregator, int64, error)
}
