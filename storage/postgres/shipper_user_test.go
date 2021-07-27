package postgres_test

import (
	pb "genproto/user_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createShipperUser(t *testing.T) *pb.ShipperUser {
	shipperUser := &pb.ShipperUser{
		Name:       fakeData.Name(),
		Username:   fakeData.UserName(),
		UserRoleId: CreateRandomId(t),
		Password:   fakeData.Characters(4),
		Phone:      fakeData.PhoneNumber(),
		IsBlocked:  logic(),
		ShipperId:  CreateRandomId(t),
	}

	res, err := strg.ShipperUser().Create(shipperUser)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	shipperUser.Id = res

	return shipperUser
}

func deleteShipperUser(t *testing.T, id, shipperId string) {
	err := strg.ShipperUser().Delete(id, shipperId)

	assert.NoError(t, err)
}

func TestCreateShipperUser(t *testing.T) {
	res := createShipperUser(t)

	assert.NotEmpty(t, res)
}

func TestUpdateShipperUser(t *testing.T) {
	shipperUser := createShipperUser(t)
	shipperUser.Name = fakeData.Name()
	shipperUser.Password = fakeData.Characters(4)

	err := strg.ShipperUser().Update(shipperUser)

	assert.NoError(t, err)
}

func TestGetShipperUser(t *testing.T) {
	id := createShipperUser(t).Id

	res, err := strg.ShipperUser().Get(id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetShipperUserByUserName(t *testing.T) {
	username := createShipperUser(t).Username

	res, err := strg.ShipperUser().GetByUsername(username)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllShipperUsers(t *testing.T) {
	res, _, err := strg.ShipperUser().GetAll(1, 10, "", "", "")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestShipperUserChangePassword(t *testing.T) {
	shipperUser := createShipperUser(t)

	err := strg.ShipperUser().ChangePassword(shipperUser.Id, fakeData.Characters(4))

	assert.NoError(t, err)
}
