package postgres_test

import (
	pb "genproto/user_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createJowiCredentials(t *testing.T) *pb.JowiCredentials {
	jowiCredentials := &pb.JowiCredentials{
		ShipperId:    CreateRandomId(t),
		DispatcherId: CreateRandomId(t),
	}

	res, err := strg.JowiCredentials().Create(jowiCredentials)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return jowiCredentials
}

func deleteJowiCredentials(t *testing.T, shipperId string) {
	err := strg.JowiCredentials().Delete(shipperId)

	assert.NoError(t, err)
}

func TestCreateJowiCredentials(t *testing.T) {
	res := createJowiCredentials(t)

	assert.NotEmpty(t, res)
}

func TestUpdateJowiCredentials(t *testing.T) {
	jowiC := createJowiCredentials(t)
	jowiC.DispatcherId = CreateRandomId(t)

	err := strg.JowiCredentials().Update(jowiC)

	assert.NoError(t, err)
}

func TestGetJowiCredentials(t *testing.T) {
	jowiC := createJowiCredentials(t)

	res, err := strg.JowiCredentials().Get(jowiC.ShipperId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteJowiC(t *testing.T) {
	shipperId := createJowiCredentials(t).ShipperId

	deleteJowiCredentials(t, shipperId)
}
