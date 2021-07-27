package postgres_test

import (
	"fmt"
	pb "genproto/user_service"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func createBranch(t *testing.T) *pb.Branch {
	location := &pb.Location{
		Long: fakeData.Longitude(),
		Lat:  fakeData.Latitude(),
	}
	branch := &pb.Branch{
		ShipperId:     CreateRandomId(t),
		Name:          fakeData.CompanyName(),
		Image:         wrapperspb.String(fakeData.Characters(10)),
		Destination:   wrapperspb.String(fakeData.Characters(12)),
		Phone:         fakeData.PhoneNumber(),
		Address:       fakeData.StreetAddress(),
		Location:      location,
		IikoId:        wrapperspb.String(CreateRandomId(t)),
		JowiId:        wrapperspb.String(CreateRandomId(t)),
		WorkHourStart: "2016-06-22 19:10:25",
		WorkHourEnd:   "2016-06-22 19:10:25-07",
	}

	res, err := strg.Branch().Create(branch)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	branch.Id = res

	return branch
}

func deleteBranch(t *testing.T, id, shipperId string) {
	err := strg.Branch().Delete(id, shipperId)

	assert.NoError(t, err)
}

func TestCreateBranch(t *testing.T) {
	res := createBranch(t)

	assert.NotEmpty(t, res)
}

func TestUpdateBranch(t *testing.T) {
	branch := createBranch(t)

	branch.Name = fakeData.CompanyName()
	branch.Phone = fakeData.PhoneNumber()

	err := strg.Branch().Update(branch)

	assert.NoError(t, err)
}

func TestGetBranch(t *testing.T) {
	phone := createBranch(t).Phone

	res, err := strg.Branch().Get(phone)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllBranches(t *testing.T) {
	shipeprId := createBranch(t).ShipperId
	res, count, err := strg.Branch().GetAll(shipeprId, "", "", 1, 10, false, false)

	assert.NoError(t, err)

	fmt.Println(res, count)
}

func TestGetByNameBranch(t *testing.T) {
	res := createBranch(t)
	shipperId := res.ShipperId
	name := res.Name

	resp, err := strg.Branch().GetByName(shipperId, name)

	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetByJowiId(t *testing.T) {
	branch := createBranch(t)
	jowiId := branch.JowiId.String()

	res, err := strg.Branch().GetByJowiID(jowiId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetByIikoId(t *testing.T) {
	branch := createBranch(t)
	iikoId := branch.IikoId.String()
	res, err := strg.Branch().GetByIikoID(iikoId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDelteBranch(t *testing.T) {
	res := createBranch(t)
	id := res.Id
	shipperId := res.ShipperId

	deleteBranch(t, id, shipperId)
}
