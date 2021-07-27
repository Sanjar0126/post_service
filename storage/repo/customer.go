package repo

import (
	pb "genproto/user_service"

	_ "github.com/lib/pq"
)

type CustomerStorageI interface {
	Create(customer *pb.Customer) (string, error)
	Get(id string) (*pb.Customer, error)
	GetAll(shipperId, search, customerTypeId string, page, limit uint64) ([]*pb.Customer, uint64, error)
	GetAggregate(shipperId, search string, page, limit uint64) ([]*pb.Customer, uint64, error)
	GetNonAggregate(shipperId, search string, page, limit uint64) ([]*pb.Customer, uint64, error)
	Update(customer *pb.Customer) error
	Delete(shipperId, id string) error
	SearchByPhone(shipperId, phone string, limit uint64, customerTypeId string) ([]*pb.Customer, error)
	GetOrInsert(shipperId, phone, name string) (string, error)
	GetByPhone(shipperId, phone string) (*pb.Customer, error)
	UpdateFcmToken(id, shipperId, fcmToken, platformID string) error
	UpdateTgChatID(id, shipperId, tgChatID string) error
	GetRegisteredCustomersReport(shipperId string, month, year uint64) ([]*pb.MontlyRegisteredCustomersReport, error)
	AttachBotLanguage(id, lang string) error
}
