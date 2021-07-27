package postgres_test

import (
	pb "genproto/user_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createPaymeInfo(t *testing.T) *pb.Payme {
	payme := &pb.Payme{
		ShipperId:  CreateRandomId(t),
		MerchantId: CreateRandomId(t),
		Login:      fakeData.FirstName(),
		Key:        fakeData.City(),
		BranchId:   CreateRandomId(t),
		BranchName: fakeData.CompanyName(),
	}

	res, err := strg.PaymeInfo().Create(payme)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return payme
}

func deletePaymeInfo(t *testing.T, id, branchId string) {
	err := strg.PaymeInfo().Delete(id, branchId)

	assert.NoError(t, err)
}

func TestCreatePaymeInfo(t *testing.T) {
	payme := createPaymeInfo(t)

	assert.NotEmpty(t, payme)
}

func TestUpdatePaymeInfo(t *testing.T) {
	payme := createPaymeInfo(t)

	payme.BranchName = fakeData.CompanyName()
	payme.Login = fakeData.FirstName()

	err := strg.PaymeInfo().Update(payme)

	assert.NoError(t, err)
}

func TestGetPaymeInfo(t *testing.T) {
	payme := createPaymeInfo(t)

	res, err := strg.PaymeInfo().Get(payme.ShipperId, "")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllPaymeInfo(t *testing.T) {
	shipperId := createPaymeInfo(t).ShipperId
	res, _, err := strg.PaymeInfo().GetAll(shipperId, nil, 1, 10)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeletePaymeInfo(t *testing.T) {
	payme := createPaymeInfo(t)

	deletePaymeInfo(t, payme.ShipperId, payme.BranchId)
}

func TestGetShipperByCredentials(t *testing.T) {
	_ = createPaymeInfo(t)

	shipperId, err := strg.PaymeInfo().GetShipperByCredentials("S2V2ZW46Q29ya2VyeWxhbmQ=") //change it

	assert.NotEmpty(t, shipperId)
	assert.NoError(t, err)
}
