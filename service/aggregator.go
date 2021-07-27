package service

import (
	"context"

	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"
	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

// AggregatorService ...
type AggregatorService struct {
	storage storage.StorageI
	logger  l.Logger
}

// NewAggregatorService ...
func NewAggregatorService(strg storage.StorageI, log l.Logger) *AggregatorService {
	return &AggregatorService{
		storage: strg,
		logger:  log,
	}
}

// Create is function for creating a courier
func (s *AggregatorService) Create(ctx context.Context, req *pb.Aggregator) (*pb.AggregatorId, error) {
	aggregatorID, err := s.storage.Aggregator().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating aggregator", req)
	}

	return &pb.AggregatorId{
		Id: aggregatorID,
	}, nil
}

// Get is function for getting a aggregator
func (s *AggregatorService) Get(ctx context.Context, req *pb.AggregatorId) (*pb.Aggregator, error) {
	aggregator, err := s.storage.Aggregator().Get(req.GetId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting aggregator", req)
	}

	return aggregator, nil
}

// GetAll is function for getting all couriers
func (s *AggregatorService) GetAll(ctx context.Context, req *pb.GetAllAggregatorsRequest) (*pb.GetAllAggregatorsResponse, error) {
	aggregators, count, err := s.storage.Aggregator().GetAll(req.Page, req.Limit, req.GetName(), req.GetShipperId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all aggregator", req)
	}

	return &pb.GetAllAggregatorsResponse{
		Aggregators: aggregators,
		Count:       uint64(count),
	}, nil
}

// Update is function for updating a aggregator
func (s *AggregatorService) Update(ctx context.Context, req *pb.Aggregator) (*gpb.Empty, error) {
	err := s.storage.Aggregator().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating aggregator", req)
	}

	return &gpb.Empty{}, nil
}

//Delete if function for deleting aggregator
func (s *AggregatorService) Delete(ctx context.Context, req *pb.AggregatorId) (*gpb.Empty, error) {
	err := s.storage.Aggregator().Delete(req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting aggregator", req)
	}

	return &gpb.Empty{}, nil
}
