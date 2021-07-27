package service

import (
	"context"

	pb "genproto/user_service"

	gpb "github.com/golang/protobuf/ptypes/empty"

	l "gitlab.udevs.io/delever/delever_user_service/pkg/logger"
	"gitlab.udevs.io/delever/delever_user_service/storage"
)

type TgBotService struct {
	storage storage.StorageI
	logger  l.Logger
}

func NewTgBotService(strg storage.StorageI, log l.Logger) *TgBotService {
	return &TgBotService{
		storage: strg,
		logger:  log,
	}
}

func (t *TgBotService) Create(ctx context.Context, req *pb.TgBot) (*pb.ShipperId, error) {
	shipperId, err := t.storage.TgBot().Create(req)

	if err != nil {
		return nil, handleError(t.logger, err, "error while creating tgbot", req)
	}

	return &pb.ShipperId{Id: shipperId}, nil
}

func (t *TgBotService) Update(ctx context.Context, req *pb.TgBot) (*gpb.Empty, error) {
	err := t.storage.TgBot().Update(req)

	if err != nil {
		return nil, handleError(t.logger, err, "error while updating tgbot", req)
	}

	return &gpb.Empty{}, nil
}

func (t *TgBotService) Delete(ctx context.Context, req *pb.ShipperId) (*gpb.Empty, error) {
	err := t.storage.TgBot().Delete(req)

	if err != nil {
		return nil, handleError(t.logger, err, "error while deleting tgbot", req)
	}

	return &gpb.Empty{}, nil
}

func (t *TgBotService) Get(ctx context.Context, req *pb.ShipperId) (*pb.TgBot, error) {
	tgbot, err := t.storage.TgBot().Get(req.GetId())

	if err != nil {
		return nil, handleError(t.logger, err, "error while getting tgbot", req)
	}

	return tgbot, nil
}

func (t *TgBotService) GetAll(ctx context.Context, req *pb.GetAllTgBotsRequest) (*pb.GetAllTgBotsResponse, error) {
	tgBots, count, err := t.storage.TgBot().GetAll(req.GetPage(), req.GetLimit())

	if err != nil {
		return nil, handleError(t.logger, err, "error while getting all tgbots", req)
	}

	return &pb.GetAllTgBotsResponse{
		TgBots: tgBots,
		Count:  count,
	}, nil
}
