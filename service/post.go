package service

import (
	"context"
	"database/sql"
	"fmt"
	pb "genproto/post_service"
	gpb "github.com/golang/protobuf/ptypes/empty"
	"gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostService struct {
	storage storage.StorageI
	logger logger.Logger
}

func NewPostService(strg storage.StorageI, log logger.Logger) *PostService {
	return &PostService{
		storage: strg,
		logger: log,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.Post) (*pb.PostId, error){
	postId, err := s.storage.Post().Create(req)
	if err != nil{
		return nil, handleError(s.logger, err, "error while creating new post", req)
	}

	return &pb.PostId{
		Id: postId,
	}, nil
}

func (s *PostService) Get(ctx context.Context, req *pb.PostId) (*pb.Post, error) {
	post, err := s.storage.Post().Get(req.GetId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting post", req)
	}

	return post, nil
}

func (s *PostService) GetAll(ctx context.Context, req *pb.GetAllPostsRequest)  (*pb.GetAllPostsResponse, error) {
	posts, count, err := s.storage.Post().GetAll(req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all posts", req)
	}

	return &pb.GetAllPostsResponse{
		Posts: posts,
		Count: count,
	}, nil
}

func (s *PostService) Update(ctx context.Context, req *pb.Post) (*gpb.Empty, error) {
	err := s.storage.Post().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating post", req)
	}
	return &gpb.Empty{}, nil
}

func (s *PostService) Delete(ctx context.Context, req *pb.DeletePostRequest) (*gpb.Empty, error) {
	err := s.storage.Post().Delete(req.GetId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting post", req)
	}

	return &gpb.Empty{}, nil
}

func handleError(log logger.Logger, err error, msg string, req interface{}) error {
	if err == sql.ErrNoRows {
		log.Error(fmt.Sprintf("%s, Not found", msg), logger.Any("req", req))
		return status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		log.Error(msg, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, "Internal server error")
	}

	return nil
}
