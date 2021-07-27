package postgres_test

import (
	pb "genproto/user_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createIikoCredentials(t *testing.T) *pb.IikoCredentials {
	iikoCredentials := &pb.IikoCredentials{
		ShipperId:    CreateRandomId(t),
		ApiLogin:     fakeData.Characters(6),
		DispatcherId: CreateRandomId(t),
	}

	res, err := strg.IikoCredentials().Create(iikoCredentials)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return iikoCredentials
}

func deleteIikoCredentials(t *testing.T, shipperId string) {
	err := strg.IikoCredentials().Delete(shipperId)

	assert.NoError(t, err)
}

func TestCreateIikoCredentials(t *testing.T) {
	res := createIikoCredentials(t)

	assert.NotEmpty(t, res)
}

func TestUpdateIikoCredentials(t *testing.T) {
	res := createIikoCredentials(t)

	res.ApiLogin = fakeData.Characters(5)
	res.DispatcherId = CreateRandomId(t)

	err := strg.IikoCredentials().Update(res)

	assert.NoError(t, err)
}

func TestGetIikoCredentials(t *testing.T) {
	shipperId := createIikoCredentials(t).ShipperId

	iikoCredentials, err := strg.IikoCredentials().Get(shipperId)

	assert.NoError(t, err)
	assert.NotEmpty(t, iikoCredentials)
}

func TestDeleteIikoCredentials(t *testing.T) {
	shipperId := createIikoCredentials(t).ShipperId

	deleteIikoCredentials(t, shipperId)
}
