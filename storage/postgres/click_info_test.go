package postgres_test

import (
	pb "genproto/user_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createClickInfo(t *testing.T) *pb.Click {
	click := &pb.Click{
		ShipperId:      CreateRandomId(t),
		MerchantId:     int64(fakeData.Rand.Intn(5000)),
		ServiceId:      int64(fakeData.Rand.Intn(5000)),
		Key:            fakeData.Characters(6),
		BranchId:       CreateRandomId(t),
		MerchantUserId: int64(fakeData.Rand.Intn(5000)),
	}

	res, err := strg.ClickInfo().Create(click)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return click
}

func deleteClickInfo(t *testing.T, shipperId, branchId string) {
	err := strg.ClickInfo().Delete(shipperId, branchId)

	assert.NoError(t, err)
}

func TestCreateClick(t *testing.T) {
	res := createClickInfo(t)

	assert.NotEmpty(t, res)
}

func TestUpdateClickInfo(t *testing.T) {
	clickInfo := createClickInfo(t)

	clickInfo.Key = fakeData.Characters(6)

	err := strg.ClickInfo().Update(clickInfo)

	assert.NoError(t, err)
}

func TestGetClickInfo(t *testing.T) {
	clickInfo := createClickInfo(t)

	res, err := strg.ClickInfo().Get(clickInfo.ShipperId, clickInfo.BranchId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllClickInfo(t *testing.T) {
	shipperId := createClickInfo(t).ShipperId
	res, _, err := strg.ClickInfo().GetAll(shipperId, nil, 1, 10)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetShipperAndKeyByCredentials(t *testing.T) {
	clickInfo := createClickInfo(t)

	_, _, err := strg.ClickInfo().GetShipperAndKeyByCredentials(clickInfo.ServiceId)

	assert.NoError(t, err)
}

func TestDeleteClickInfo(t *testing.T) {
	res := createClickInfo(t)

	deleteClickInfo(t, res.ShipperId, res.BranchId)
}
