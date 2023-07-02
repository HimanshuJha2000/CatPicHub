package service

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/model"
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

var svc service.Service

func TestGetCatPicByIDService(t *testing.T) {
	// Mock data
	catPicId := "123"

	// Patch the GetByID function to return a mock CatPics object
	defer monkey.PatchInstanceMethod(reflect.TypeOf(&model.CatPics{}), "GetByID",
		func(cpo *model.CatPics, catPicId string) error {
			cpo.ID = "image123"
			cpo.CatPicFileName = "cat.jpg"
			cpo.CatPicData = []byte{0x12, 0x34, 0x56}
			return nil
		}).Unpatch()

	// Call the function under test
	statusCode, _, _ := svc.GetCatPicByIDService(catPicId)
	assert.Equal(t, statusCode, http.StatusOK)
}

func TestGetCatPicByIDServiceError(t *testing.T) {
	// Mock data
	catPicId := "123"

	// Patch the GetByID function to return a mock CatPics object
	defer monkey.PatchInstanceMethod(reflect.TypeOf(&model.CatPics{}), "GetByID",
		func(cpo *model.CatPics, catPicId string) error {
			return fmt.Errorf("error")
		}).Unpatch()

	// Call the function under test
	statusCode, _, _ := svc.GetCatPicByIDService(catPicId)
	assert.Equal(t, statusCode, http.StatusBadRequest)
}

func TestGetCatPicByIDServicePngImage(t *testing.T) {
	// Mock data
	catPicId := "123"

	// Patch the GetByID function to return a mock CatPics object
	defer monkey.PatchInstanceMethod(reflect.TypeOf(&model.CatPics{}), "GetByID",
		func(cpo *model.CatPics, catPicId string) error {
			cpo.ID = "image123"
			cpo.CatPicFileName = "cat.png"
			return nil
		}).Unpatch()

	// Call the function under test
	statusCode, _, _ := svc.GetCatPicByIDService(catPicId)
	assert.Equal(t, statusCode, http.StatusOK)
}

func TestGetCatPicsListService(t *testing.T) {
	// Mock data
	pageNo := 1
	pageSize := 10
	expectedImages := []model.CatPics{
		{ID: "image1", CatPicFileName: "cat1.jpg", CatPicData: []byte{0x12, 0x34, 0x56}},
		{ID: "image2", CatPicFileName: "cat2.jpg", CatPicData: []byte{0x78, 0x9A, 0xBC}},
	}

	// Patch the GetPaginatedData function to return mock data
	defer monkey.Patch(model.GetPaginatedData,
		func(pageNo, pageSize int) ([]model.CatPics, error) {
			return expectedImages, nil
		}).Unpatch()

	_, result, _ := svc.GetCatPicsListService(pageNo, pageSize)
	assert.Equal(t, result["images"], expectedImages)
}

func TestGetCatPicsListServiceError(t *testing.T) {
	// Mock data
	pageNo := 1
	pageSize := 10

	// Patch the GetPaginatedData function to return mock data
	defer monkey.Patch(model.GetPaginatedData,
		func(pageNo, pageSize int) ([]model.CatPics, error) {
			return nil, fmt.Errorf("error")
		}).Unpatch()

	statusCode, _, _ := svc.GetCatPicsListService(pageNo, pageSize)
	assert.Equal(t, statusCode, http.StatusBadRequest)
}

func TestDeleteCatPicServiceErrorInGet(t *testing.T) {
	// Mock data
	catPicID := "123"

	// Patch the GetByID function to return a mock CatPics object
	patchGetByID := monkey.PatchInstanceMethod(reflect.TypeOf(&model.CatPics{}), "GetByID",
		func(cpo *model.CatPics, catPicId string) error {
			return fmt.Errorf("error")
		})
	defer patchGetByID.Unpatch()

	statusCode, _, _ := svc.DeleteCatPicService(catPicID)
	assert.Equal(t, statusCode, http.StatusBadRequest)
}
