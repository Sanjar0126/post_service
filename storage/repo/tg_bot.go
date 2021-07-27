package repo

import (
	pb "genproto/user_service"
)

type TgBotsStorageI interface {
	Create(tg *pb.TgBot) (string, error)
	Update(tg *pb.TgBot) error
	Delete(id *pb.ShipperId) error
	Get(id string) (*pb.TgBot, error)
	GetAll(page, limit int64) ([]*pb.TgBot, int64, error)
}
