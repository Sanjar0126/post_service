package postgres_test

import (
	pb "genproto/user_service"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func createBranchUser(t *testing.T) *pb.BranchUser {
	branchUser := &pb.BranchUser{
		Name:       fakeData.Name(),
		Phone:      fakeData.PhoneNumber(),
		IsActive:   logic(),
		IsBlocked:  logic(),
		ShipperId:  CreateRandomId(t),
		UserRoleId: CreateRandomId(t),
		FcmToken:   wrapperspb.String(fakeData.Characters(12)),
		BranchId:   CreateRandomId(t),
	}

	res, err := strg.BranchUser().Create(branchUser)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	branchUser.Id = res

	return branchUser
}

func deleteBranchUser(t *testing.T, id, shipperId string) {
	err := strg.BranchUser().Delete(id, shipperId)

	assert.NoError(t, err)
}

func TestCreateBranchUser(t *testing.T) {
	res := createBranchUser(t)

	assert.NotEmpty(t, res)
}

func TestUpdateBranchUser(t *testing.T) {
	res := createBranchUser(t)
	res.Name = fakeData.Name()
	res.Phone = fakeData.PhoneNumber()

	err := strg.BranchUser().Update(res)

	assert.NoError(t, err)
}

func TestGetBranchUser(t *testing.T) {
	id := createBranchUser(t).Id

	res, err := strg.BranchUser().Get(id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllBranchUsers(t *testing.T) {
	shipperId := createBranchUser(t).ShipperId
	res, _, err := strg.BranchUser().GetAll(1, 10, shipperId, "", "", "")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUpdateFCMToken(t *testing.T) {
	res := createBranchUser(t)
	id := res.Id
	shipperId := res.ShipperId
	fcmToken := fakeData.Characters(12)

	err := strg.BranchUser().UpdateFcmToken(id, shipperId, fcmToken, CreateRandomId(t))

	assert.NoError(t, err)
}

func TestDeleteFCMToken(t *testing.T) {
	id := createBranchUser(t).Id

	err := strg.BranchUser().DeleteFcmToken(id)

	assert.NoError(t, err)
}

func TestGetByPhone(t *testing.T) {
	res := createBranchUser(t)
	phone := res.Phone
	shipperId := res.ShipperId

	get, err := strg.BranchUser().GetByPhone(phone, shipperId)

	assert.NoError(t, err)
	assert.NotEmpty(t, get)
}

func TestDeleteBranchUser(t *testing.T) {
	res := createBranchUser(t)

	deleteBranchUser(t, res.Id, res.ShipperId)
}
