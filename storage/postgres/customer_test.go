package postgres_test

import (
	pb "genproto/user_service"
	"math/rand"
	"testing"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createCustomer(t *testing.T) *pb.Customer {
	customer := &pb.Customer{
		Name:        "fakeData.Name()",
		Phone:       fakeData.PhoneNumber(),
		IsBlocked:   1 == 2,
		ShipperId:   CreateRandomId(t),
		TgChatId:    wrapperspb.String(CreateRandomId(t)),
		FcmToken:    wrapperspb.String(CreateRandomId(t)),
		DateOfBirth: wrapperspb.String("2020-04-04"),
		IsAggregate: rand.Intn(2) == 1,
	}

	res, err := strg.Customer().Create(customer)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	customer.Id = res

	return customer
}

func deleteCustomer(t *testing.T, shipperId, id string) {
	err := strg.Customer().Delete(shipperId, id)

	assert.NoError(t, err)
}

func TestCreateCustomer(t *testing.T) {
	res := createCustomer(t)

	assert.NotEmpty(t, res)
}

func TestUpdateCustomer(t *testing.T) {
	customer := createCustomer(t)
	customer.Name = fakeData.Name()
	customer.Phone = fakeData.PhoneNumber()
	customer.IsBlocked = true
	customer.DateOfBirth = wrapperspb.String("2020-04-04")
	err := strg.Customer().Update(customer)

	assert.NoError(t, err)
}

func TestGetCustomer(t *testing.T) {
	id := createCustomer(t).Id

	res, err := strg.Customer().Get(id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllCustomer(t *testing.T) {
	shipperId := createCustomer(t).ShipperId

	res, count, err := strg.Customer().GetAll(shipperId, "", "", 1, 10)
	assert.NoError(t, err)

	if len(res) != int(count) {
		assert.NoError(t, err)
	}

	assert.NotEmpty(t, res)
}

func TestGetAggregate(t *testing.T) {
	shipperId := createCustomer(t).ShipperId

	res, count, err := strg.Customer().GetAggregate(shipperId, "", 1, 10)
	assert.NoError(t, err)

	if len(res) != int(count) {
		assert.NoError(t, err)
	}

	assert.NotEmpty(t, res)
}

func TestGetNonAggregate(t *testing.T) {
	shipperId := createCustomer(t).ShipperId

	res, count, err := strg.Customer().GetNonAggregate(shipperId, "", 1, 10)
	assert.NoError(t, err)

	if len(res) != int(count) {
		assert.NoError(t, err)
	}

	assert.NotEmpty(t, res)
}

func TestCustomerSearchByPhone(t *testing.T) {
	customer := createCustomer(t)
	phone := customer.Phone
	shipperId := customer.ShipperId
	res, err := strg.Customer().SearchByPhone(shipperId, phone, 10, "")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetOrInsertCustomer(t *testing.T) {
}

func TestGetByPhoneCustomer(t *testing.T) {
	customer := createCustomer(t)
	phone := customer.Phone
	shipperId := customer.ShipperId

	res, err := strg.Customer().GetByPhone(shipperId, phone)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUpdateFCMTokenCustomer(t *testing.T) {
	uuid, _ := uuid.NewUUID()

	customer := createCustomer(t)
	id := customer.Id
	shipperId := customer.ShipperId
	fcmToken := uuid.String()
	platformID := uuid.String()

	err := strg.Customer().UpdateFcmToken(id, shipperId, fcmToken, platformID)

	assert.NoError(t, err)
}

func TestUpdateTgCHatId(t *testing.T) {
	customer := createCustomer(t)
	tgChatId := CreateRandomId(t)
	id := customer.Id
	shipperId := customer.ShipperId

	err := strg.Customer().UpdateTgChatID(id, shipperId, tgChatId)

	assert.NoError(t, err)
}

func TestAttachBotLanguage(t *testing.T) {
	id := createCustomer(t).Id
	language := fakeData.Language

	err := strg.Customer().AttachBotLanguage(id, language)

	assert.NoError(t, err)
}

func TestDeleteCustomer(t *testing.T) {
	res := createCustomer(t)
	id := res.Id
	shipperId := res.ShipperId
	deleteCustomer(t, shipperId, id)
}
