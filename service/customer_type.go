package service

import (
	"context"
	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"
	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

type CustomerTypeService struct {
	storage storage.StorageI
	logger  l.Logger
}

func NewCustomerTypeService(strg storage.StorageI, log l.Logger) *CustomerTypeService {
	return &CustomerTypeService{
		storage: strg,
		logger:  log,
	}
}

// Create is function for creating a customer_type
func (s *CustomerTypeService) Create(ctx context.Context, req *pb.CustomerType) (*pb.CustomerTypeId, error) {
	customer_type_id, err := s.storage.CustomerType().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating customer_type", req)
	}

	return &pb.CustomerTypeId{
		Id: customer_type_id,
	}, nil
}

// Get is function for getting a customer_type
func (s *CustomerTypeService) Get(ctx context.Context, req *pb.CustomerTypeId) (*pb.CustomerType, error) {
	customer_type, err := s.storage.CustomerType().Get(req.GetId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting customer_type", req)
	}

	return customer_type, nil
}

// GetAll is function for getting all couriers
func (s *CustomerTypeService) GetAll(ctx context.Context, req *pb.GetAllCustomerTypeRequest) (*pb.GetAllCustomerTypeResponse, error) {
	customer_types, count, err := s.storage.CustomerType().GetAll(req.Page, req.Limit, req.GetName())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all customer_type", req)
	}

	return &pb.GetAllCustomerTypeResponse{
		CustomerTypes: customer_types,
		Count:         uint64(count),
	}, nil
}

// Update is function for updating a customer_type
func (s *CustomerTypeService) Update(ctx context.Context, req *pb.CustomerType) (*gpb.Empty, error) {
	err := s.storage.CustomerType().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating customer_type", req)
	}

	return &gpb.Empty{}, nil
}

// Delete if function for deleting customer_type
func (s *CustomerTypeService) Delete(ctx context.Context, req *pb.CustomerTypeId) (*gpb.Empty, error) {
	err := s.storage.CustomerType().Delete(req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting customer_type", req)
	}

	return &gpb.Empty{}, nil
}
