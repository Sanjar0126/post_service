package service

import (
	"context"
	"genproto/order_service"
	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"gitlab.udevs.io/delever/delever_user_service/pkg/grpc_client"
	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

// CustomerService ...
type CustomerService struct {
	storage storage.StorageI
	logger  l.Logger
	client  *grpc_client.GrpcClient
}

// NewCustomerService ...
func NewCustomerService(strg storage.StorageI, log l.Logger, client *grpc_client.GrpcClient) *CustomerService {
	return &CustomerService{
		storage: strg,
		logger:  log,
		client:  client,
	}
}

// Create is function for creating a courier
func (s *CustomerService) Create(ctx context.Context, req *pb.Customer) (*pb.CustomerId, error) {
	customerID, err := s.storage.Customer().Create(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while creating customer", req)
	}

	go func() {
		_, err = s.client.CustomerOrderService().Create(
			context.Background(),
			&order_service.CustomerOrderService{
				Id:             customerID,
				Name:           req.Name,
				Phone:          req.Phone,
				ShipperId:      req.ShipperId,
				DateOfBirth:    req.DateOfBirth,
				CustomerTypeId: req.CustomerTypeId,
			},
		)

		if err != nil {
			handleError(s.logger, err, "error while creating customer on order service", req)
		}
	}()

	return &pb.CustomerId{
		Id: customerID,
	}, nil
}

// Update is function for updating a customer
func (s *CustomerService) Update(ctx context.Context, req *pb.Customer) (*gpb.Empty, error) {
	err := s.storage.Customer().Update(req)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating customer", req)
	}

	go func() {
		_, err = s.client.CustomerOrderService().Update(
			context.Background(),
			&order_service.CustomerOrderService{
				Id:             req.Id,
				Name:           req.Name,
				Phone:          req.Phone,
				ShipperId:      req.ShipperId,
				DateOfBirth:    req.DateOfBirth,
				CustomerTypeId: req.CustomerTypeId,
			},
		)

		if err != nil {
			handleError(s.logger, err, "error while updating customer on order service", req)
		}
	}()

	return &gpb.Empty{}, nil
}

// Delete if function for deleting customer
func (s *CustomerService) Delete(ctx context.Context, req *pb.DeleteCustomerRequest) (*gpb.Empty, error) {
	err := s.storage.Customer().Delete(req.ShipperId, req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while deleting customer", req)
	}

	return &gpb.Empty{}, nil
}

// Get is function for getting a customer
func (s *CustomerService) Get(ctx context.Context, req *pb.CustomerId) (*pb.Customer, error) {
	customer, err := s.storage.Customer().Get(req.Id)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting customer", req)
	}

	return customer, nil
}

// GetAll is function for getting all couriers
func (s *CustomerService) GetAll(ctx context.Context, req *pb.GetAllCustomersRequest) (*pb.GetAllCustomersResponse, error) {
	customers, count, err := s.storage.Customer().GetAll(req.ShipperId, req.GetSearch(), req.CustomerTypeId, req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all customers", req)
	}

	return &pb.GetAllCustomersResponse{
		Customers: customers,
		Count:     count,
	}, nil
}

// GetAggregate is function for getting all aggregate customers
func (s *CustomerService) GetAggregate(ctx context.Context, req *pb.GetAllCustomersRequest) (*pb.GetAllCustomersResponse, error) {
	customers, count, err := s.storage.Customer().GetAggregate(req.ShipperId, req.GetSearch(), req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all aggregate customers", req)
	}

	return &pb.GetAllCustomersResponse{
		Customers: customers,
		Count:     count,
	}, nil
}

// GetNonAggregate is function for getting all non aggregate customers
func (s *CustomerService) GetNonAggregate(ctx context.Context, req *pb.GetAllCustomersRequest) (*pb.GetAllCustomersResponse, error) {
	customers, count, err := s.storage.Customer().GetNonAggregate(req.ShipperId, req.GetSearch(), req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting all non aggregate customers", req)
	}

	return &pb.GetAllCustomersResponse{
		Customers: customers,
		Count:     count,
	}, nil
}

// SearchByPhone ....
func (s *CustomerService) SearchByPhone(ctx context.Context, req *pb.SearchCustomersByPhoneRequest) (*pb.SearchCustomersByPhoneResponse, error) {
	customers, err := s.storage.Customer().SearchByPhone(req.ShipperId, req.Phone, req.GetLimit(), req.GetCustomerTypeId())
	if err != nil {
		return nil, handleError(s.logger, err, "error while searching customer by phone", req)
	}

	return &pb.SearchCustomersByPhoneResponse{
		Customers: customers,
	}, nil
}

// GetOrInsert is a function to get or insert function
func (s *CustomerService) GetOrInsert(ctx context.Context, req *pb.GetCustomerOrInsertRequest) (*pb.CustomerId, error) {
	customerID, err := s.storage.Customer().GetOrInsert(req.ShipperId, req.Phone, req.Name)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting or inserting customer", req)
	}

	go func() {
		_, err = s.client.CustomerOrderService().GetOrInsert(
			context.Background(),
			&order_service.CustomerOrderService{
				Id:        customerID,
				Name:      req.Name,
				Phone:     req.Phone,
				ShipperId: req.ShipperId,
			},
		)

		if err != nil {
			handleError(s.logger, err, "error while updating customer on order service", req)
		}
	}()

	return &pb.CustomerId{
		Id: customerID,
	}, nil
}

// GetByPhone is a function to get a customer by phone
func (s *CustomerService) GetByPhone(ctx context.Context, req *pb.GetCustomerByPhoneRequest) (*pb.Customer, error) {
	customer, err := s.storage.Customer().GetByPhone(req.ShipperId, req.Phone)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting customer by phone", req)
	}

	return customer, nil
}

// UpdateFcmToken is function to update Fcm token
func (s *CustomerService) UpdateFcmToken(ctx context.Context, req *pb.UpdateCustomerFcmTokenRequest) (*gpb.Empty, error) {
	err := s.storage.Customer().UpdateFcmToken(req.Id, req.ShipperId, req.FcmToken, req.PlatformId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating customer FcmToken", req)
	}

	go func() {
		_, err = s.client.CustomerOrderService().UpdatePlatform(
			context.Background(),
			&order_service.CustomerOrderService{
				Id:        req.Id,
				ShipperId: req.ShipperId,
				PlatformId: &wrapperspb.StringValue{
					Value: req.PlatformId,
				},
			},
		)

		if err != nil {
			handleError(s.logger, err, "error while updating customer platform_id on order service", req)
		}
	}()

	return &gpb.Empty{}, nil
}

func (s *CustomerService) UpdateTgChatId(ctx context.Context, req *pb.UpdateCustomerTgChatIdRequest) (*gpb.Empty, error) {
	err := s.storage.Customer().UpdateTgChatID(req.Id, req.ShipperId, req.TgChatId)
	if err != nil {
		return nil, handleError(s.logger, err, "error while updating customer FcmToken", req)
	}

	go func() {
		_, err = s.client.CustomerOrderService().UpdateTgChatId(
			context.Background(),
			&order_service.CustomerOrderService{
				Id:        req.Id,
				ShipperId: req.ShipperId,
				TgChatId:  req.TgChatId,
			},
		)

		if err != nil {
			handleError(s.logger, err, "error while updating customer tg chat id on order service", req)
		}
	}()

	return &gpb.Empty{}, nil
}

func (s *CustomerService) GetRegisteredCustomersReport(ctx context.Context, req *pb.GetRegisteredCustomersReportRequest) (*pb.GetRegisteredCustomersReportResponse, error) {
	monthlyReports, err := s.storage.Customer().GetRegisteredCustomersReport(req.ShipperId, req.Year, req.Month)
	if err != nil {
		return nil, handleError(s.logger, err, "error while getting registered customers report", req)
	}

	return &pb.GetRegisteredCustomersReportResponse{
		Report: monthlyReports,
	}, nil
}

// AttachBotLanguage is function for attaching language to customer
func (s *CustomerService) AttachBotLanguage(ctx context.Context, req *pb.AttachBotLanguageRequest) (*gpb.Empty, error) {
	err := s.storage.Customer().AttachBotLanguage(req.Id, req.Lang)
	if err != nil {
		return nil, handleError(s.logger, err, "error while attaching language to customer", req.Id)
	}

	return &gpb.Empty{}, nil
}
