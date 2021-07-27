package postgres_test

import (
	pb "genproto/user_service"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func createShipper(t *testing.T) *pb.Shipper {
	phones := []string{fakeData.PhoneNumber(), fakeData.PhoneNumber()}
	shipper := &pb.Shipper{
		Name:                     fakeData.Name(),
		Phone:                    phones,
		Logo:                     fakeData.Characters(8),
		Description:              fakeData.Characters(12),
		IsActive:                 logic(),
		WorkHourStart:            "19:10:25",
		WorkHourEnd:              "21:10:25-07",
		CourierAcceptsFirst:      logic(),
		CheckCourierActionRadius: logic(),
		MaxCourierOrders:         int64(fakeData.Rand.Intn(20)),
		MenuImage:                wrapperspb.String(fakeData.Characters(4)),
		MaxDeliveryTime:          int64(fakeData.Rand.Intn(20)),
		Crm:                      "iiko",
		CallCenterTg:             wrapperspb.String(fakeData.CellPhoneNumber()),
	}

	res, err := strg.Shipper().Create(shipper)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	shipper.Id = res

	return shipper
}

func deleteShipper(t *testing.T, id string) {
	err := strg.Shipper().Delete(id)

	assert.NoError(t, err)
}

func TestCreateShipper(t *testing.T) {
	id := createShipper(t)

	assert.NotEmpty(t, id)
}

func TestUpdateShipper(t *testing.T) {
	shipper := createShipper(t)
	shipper.Name = fakeData.Name()
	shipper.Description = fakeData.Characters(12)
	shipper.IsActive = logic()

	err := strg.Shipper().Update(shipper)

	assert.NoError(t, err)
}

func TestGetShipper(t *testing.T) {
	id := createShipper(t).Id

	res, err := strg.Shipper().Get(id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllShippers(t *testing.T) {
	res, _, err := strg.Shipper().GetAll(1, 10, logic())

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetByName(t *testing.T) {
	name := createShipper(t).Name

	res, err := strg.Shipper().GetByName(name)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteShipper(t *testing.T) {
	id := createShipper(t).Id
	deleteShipper(t, id)
}

func logic() bool {
	return rand.Intn(2) != 0
}
