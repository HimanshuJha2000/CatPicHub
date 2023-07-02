package service

import (
	"fmt"
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/model"
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/utils"
	"mime/multipart"
	"net/http"
	"strings"
)

type Service struct{}

func (svc Service) CreateCatPicService(file *multipart.FileHeader) (int, map[string]interface{}, error) {
	// Open the file
	src, err := file.Open()
	fmt.Print(17)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status_code":    http.StatusInternalServerError,
			"external_error": "Failed to open image file",
			"internal error": err.Error(),
		}, err
	}
	defer src.Close()

	// Read the image file into a byte array
	buffer := make([]byte, file.Size)
	_, err = src.Read(buffer)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status_code":    http.StatusInternalServerError,
			"external_error": "Failed to read image file",
			"internal error": err.Error(),
		}, err

	}

	picFileType := "jpg"
	if strings.HasSuffix(file.Filename, ".png") {
		picFileType = "png"
	}

	// Save to database
	cpo := model.CatPics{
		ID:             utils.GenerateUniqueID(),
		CatPicFileName: file.Filename,
		CatPicData:     buffer,
		CatPicFileType: picFileType,
	}

	if err = cpo.Create(); err != nil {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code":    http.StatusBadRequest,
			"external_error": "Failed to save image file",
			"internal error": err.Error(),
		}, err
	}

	return http.StatusOK, map[string]interface{}{
		"image_id":        cpo.ID,
		"image_file_name": cpo.CatPicFileName,
		"image_file_type": cpo.CatPicFileType,
		"success":         "Cat picture uploaded successfully",
	}, err
}

func (svc Service) GetCatPicByIDService(catPicId string) (int, map[string]interface{}, error) {
	var cpo model.CatPics
	if err := cpo.GetByID(catPicId); err != nil {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code":    http.StatusBadRequest,
			"external_error": "Image with this ID doesn't exist",
			"internal error": err.Error(),
		}, err
	}

	// Set the appropriate content headers
	contentType := "image/jpeg"
	fileName := cpo.CatPicFileName
	if strings.HasSuffix(fileName, ".png") {
		contentType = "image/png"
	}
	return http.StatusOK, map[string]interface{}{
		"content_type":    contentType,
		"image_id":        cpo.ID,
		"image_file_name": cpo.CatPicFileName,
		"image_file_data": cpo.CatPicData,
		"success":         "Cat picture uploaded successfully",
	}, nil
}

func (svc Service) GetCatPicsListService(pageNo, pageSize int) (int, map[string]interface{}, error) {
	catPicsArray, err := model.GetPaginatedData(pageNo, pageSize)
	if err != nil {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code":    http.StatusBadRequest,
			"external_error": "Failed to read cat pics data from database",
			"internal error": err.Error(),
		}, err
	}

	return http.StatusOK, map[string]interface{}{
		"images":  catPicsArray,
		"success": "List of cat pictures fetched successfully",
	}, nil
}

func (svc Service) UpdateCatPicService(file *multipart.FileHeader, catPicID string) (int, map[string]interface{}, error) {
	var cpo model.CatPics
	if err := cpo.GetByID(catPicID); err != nil {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code":    http.StatusBadRequest,
			"external_error": "Image with this ID doesn't exist",
			"internal error": err.Error(),
		}, err
	}

	src, err := file.Open()
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status_code":    http.StatusInternalServerError,
			"external_error": "Failed to open image file",
			"internal error": err.Error(),
		}, err
	}
	defer src.Close()

	// Read the image file into a byte array
	buffer := make([]byte, file.Size)
	_, err = src.Read(buffer)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status_code":    http.StatusInternalServerError,
			"external_error": "Failed to read image file",
			"internal error": err.Error(),
		}, err
	}

	picFileType := "jpg"
	if strings.HasSuffix(file.Filename, ".png") {
		picFileType = "png"
	}

	err = cpo.Update(map[string]interface{}{
		"cat_pic_file_name": file.Filename,
		"cat_pic_file_type": picFileType,
		"cat_pic_data":      buffer,
	})
	if err != nil {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code":    http.StatusBadRequest,
			"external_error": "Failed to update the image file",
			"internal error": err.Error(),
		}, err
	}
	return http.StatusOK, map[string]interface{}{
		"image_id":        cpo.ID,
		"image_file_name": file.Filename,
		"image_file_type": picFileType,
		"success":         "Cat picture updated successfully!!",
	}, err
}

func (svc Service) DeleteCatPicService(catPicID string) (int, map[string]interface{}, error) {
	var cpo model.CatPics
	if err := cpo.GetByID(catPicID); err != nil {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code":    http.StatusBadRequest,
			"external_error": "Image with this ID doesn't exist",
			"internal error": err.Error(),
		}, err
	}
	deletedId := cpo.ID
	err := cpo.DeleteById()
	if err != nil {
		return http.StatusBadRequest, map[string]interface{}{
			"status_code":    http.StatusBadRequest,
			"external_error": "Failed to delete the image file",
			"internal error": err.Error(),
		}, err
	}

	return http.StatusOK, map[string]interface{}{
		"Deleted ImageID": deletedId,
		"success":         "Cat picture deleted successfully!!",
	}, err
}
