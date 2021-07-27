package postgres_test

import (
	pb "genproto/user_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTgBot(t *testing.T) *pb.TgBot {
	tgBot := &pb.TgBot{
		ShipperId:   CreateRandomId(t),
		BotToken:    fakeData.Characters(16),
		Url:         fakeData.URL(),
		AccessToken: fakeData.Characters(16),
	}

	res, err := strg.TgBot().Create(tgBot)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return tgBot
}

func TestCreateTgBot(t *testing.T) {
	res := createTgBot(t)

	assert.NotEmpty(t, res)
}

func TestUpdateTgBot(t *testing.T) {
	tgbot := createTgBot(t)
	tgbot.Url = fakeData.URL()
	tgbot.BotToken = fakeData.Characters(16)

	err := strg.TgBot().Update(tgbot)

	assert.NoError(t, err)
}

func TestGetTgBot(t *testing.T) {
	shipperId := createTgBot(t).ShipperId

	res, err := strg.TgBot().Get(shipperId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllTgBots(t *testing.T) {
	res, _, err := strg.TgBot().GetAll(1, 10)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
