package postgres_test

import (
	pb "genproto/user_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createSystemUser(t *testing.T) *pb.SystemUser {
	systemUser := &pb.SystemUser{
		Name:      fakeData.Name(),
		Username:  fakeData.UserName(),
		Password:  fakeData.Characters(4),
		Phone:     fakeData.PhoneNumber(),
		IsBlocked: logic(),
	}

	res, err := strg.SystemUser().Create(systemUser)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	systemUser.Id = res

	return systemUser
}

func deleteSystemUser(t *testing.T, id string) {
	err := strg.SystemUser().Delete(id)

	assert.NoError(t, err)
}

func TestCreateSystemUser(t *testing.T) {
	res := createSystemUser(t)

	assert.NotEmpty(t, res)
}

func TestUpdateSystemUser(t *testing.T) {
	systemUser := createSystemUser(t)
	systemUser.Name = fakeData.Name()
	systemUser.Password = fakeData.Characters(4)

	err := strg.SystemUser().Update(systemUser)

	assert.NoError(t, err)
}

func TestGetSystemUser(t *testing.T) {
	id := createSystemUser(t).Id

	res, err := strg.SystemUser().Get(id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllSystemUsers(t *testing.T) {
	res, _, err := strg.SystemUser().GetAll(1, 10, "")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetByUserName(t *testing.T) {
	username := createSystemUser(t).Username

	res, err := strg.SystemUser().GetByUsername(username)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestChangePassword(t *testing.T) {
	systemUser := createSystemUser(t)
	id := systemUser.Id
	password := fakeData.Characters(4)

	err := strg.SystemUser().ChangePassword(id, password)

	assert.NoError(t, err)
}
