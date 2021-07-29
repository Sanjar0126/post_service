package repo

import pb "genproto/post_service"

type PostStorageI interface {
	Create(post *pb.Post) (*pb.Post, error)
	Update(post *pb.Post) error
	Delete(id uint32) error
	Get(id uint32) (*pb.Post, error)
	GetAll(page, limit uint64) ([]*pb.Post, uint64, error)
}