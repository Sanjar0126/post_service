package service

import (
	"context"
	"database/sql"
	"fmt"
	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

// ShipperService ...
type ShipperService struct {
	storage storage.StorageI
	logger  l.Logger
}

// NewShipperService ...
func NewShipperService(strg storage.StorageI, log l.Logger) *ShipperService {
	return &ShipperService{
		storage: strg,
		logger:  log,
	}
}

// Create is function for creating a courier
func (s *ShipperService) Create(ctx context.Context, req *pb.Shipper) (*pb.ShipperId, error) {
	shipperId, err := s.storage.Shipper().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating shipper", req)
	}

	return &pb.ShipperId{
		Id: shipperId,
	}, nil
}

// Get is function for getting a branch
func (s *ShipperService) Get(ctx context.Context, req *pb.ShipperId) (*pb.Shipper, error) {
	shipper, err := s.storage.Shipper().Get(req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting shipper", req)
	}

	return shipper, nil
}

// GetAll is function for getting all shippers
func (s *ShipperService) GetAll(ctx context.Context, req *pb.GetAllShippersRequest) (*pb.GetAllShippersResponse, error) {
	shippers, count, err := s.storage.Shipper().GetAll(req.GetPage(), req.GetLimit(), req.GetHasIiko())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all shippers", req)
	}

	return &pb.GetAllShippersResponse{
		Shippers: shippers,
		Count:    count,
	}, nil
}

// Update is function for updating a branch
func (s *ShipperService) Update(ctx context.Context, req *pb.Shipper) (*gpb.Empty, error) {
	err := s.storage.Shipper().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating shipper", req)
	}

	return &gpb.Empty{}, nil
}

//Delete if function for deleting branch
func (s *ShipperService) Delete(ctx context.Context, req *pb.ShipperId) (*gpb.Empty, error) {
	err := s.storage.Shipper().Delete(req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting shipper", req)
	}

	return &gpb.Empty{}, nil
}

// GetByName is function for getting shipper by name
func (s *ShipperService) GetByName(ctx context.Context, req *pb.GetShipperByNameRequest) (*pb.Shipper, error) {
	shipper, err := s.storage.Shipper().GetByName(req.Name)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting shipper by name", req)
	}

	return shipper, nil
}

func handleError(log l.Logger, err error, msg string, req interface{}) error {
	if err == sql.ErrNoRows {
		log.Error(fmt.Sprintf("%s, Not found", msg), l.Any("req", req))
		return status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		log.Error(msg, l.Error(err), l.Any("req", req))
		return status.Error(codes.Internal, "Internal server error")
	}

	return nil
}
