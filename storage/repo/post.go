package repo

import pb "genproto/post_service"

type PostStorageI interface {
	Create(post *pb.Post) (string, error)
	Update(post *pb.Post) error
	Delete(id string) error
	Get(id string) (*pb.Post, error)
	GetAll(page, limit uint64) ([]*pb.Post, uint64, error)
}