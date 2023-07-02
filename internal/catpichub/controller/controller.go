package controller

import (
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/service"
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CatPicHubController struct {
	CatPicHubService service.Service
}

var CatPicHub CatPicHubController

func (cpc *CatPicHubController) CreateCatPicController(ctx *gin.Context) {

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           "Image file is required",
			"internal_error ": err.Error(),
		})
		return
	}

	if !strings.HasSuffix(file.Filename, ".png") && !strings.HasSuffix(file.Filename, ".jpg") && !strings.HasSuffix(file.Filename, ".jpeg") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           "Provided file is not an image file",
			"internal_error ": "Images of only .png, .jpg and .jpeg type is supported",
		})
		return
	}

	statusCode, result, err := cpc.CatPicHubService.CreateCatPicService(file)
	if err != nil {
		log.Println("Error while inserting to database is ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
	} else {
		ctx.JSON(statusCode, result)
	}
}

func (cpc *CatPicHubController) GetCatPicByIDController(ctx *gin.Context) {

	catPicID := ctx.Params.ByName("cat_pic_id")
	if catPicID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           "Failed to get Image by this ImageID",
			"internal_error ": "Cat Image ID is not provided",
		})
		return
	}

	statusCode, result, err := cpc.CatPicHubService.GetCatPicByIDService(catPicID)
	if err != nil {
		log.Println("Error occurred while fetching the cat pic from the database: ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
		return
	}

	// Write the image data as the response body
	ctx.Header("Content-Type", result["content_type"].(string))
	ctx.Writer.WriteHeader(statusCode)
	ctx.Writer.Write(result["image_file_data"].([]byte))
}

func (cpc *CatPicHubController) GetCatPicListController(ctx *gin.Context) {

	pageNo := ctx.Params.ByName("page_no")
	if pageNo == "" || utils.IsSpecialCharacter(rune(pageNo[0])) || utils.IsAlphabet(rune(pageNo[0])) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           "Page number is missing/invalid for fetching list of images",
			"internal_error ": "pageNo query param is missing/invalid!",
		})
		return
	}

	pageSize := ctx.Params.ByName("page_size")
	if pageSize == "" || utils.IsSpecialCharacter(rune(pageSize[0])) || utils.IsAlphabet(rune(pageSize[0])) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           "Page size is missing/invalid for fetching list of images",
			"internal_error ": "pageSize query param is missing/invalid!",
		})
		return
	}

	pageNoValue, _ := strconv.Atoi(pageNo)
	pageSizeValue, _ := strconv.Atoi(pageSize)
	statusCode, result, err := cpc.CatPicHubService.GetCatPicsListService(pageNoValue, pageSizeValue)

	if err != nil {
		log.Println("Error occurred while fetching the list of cat pics from the database: ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
		return
	} else {
		ctx.JSON(statusCode, result)
	}

}

func (cpc *CatPicHubController) UpdateCatPicController(ctx *gin.Context) {

	catPicID := ctx.Params.ByName("cat_pic_id")
	if catPicID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           "Failed to get Image by this ImageID",
			"internal_error ": "Cat Image ID is not provided",
		})
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           "Image file is required",
			"internal_error ": err.Error(),
		})
		return
	}

	statusCode, result, err := cpc.CatPicHubService.UpdateCatPicService(file, catPicID)
	if err != nil {
		log.Println("Error while updating image in database : ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
	} else {
		ctx.JSON(statusCode, result)
	}
}

func (cpc *CatPicHubController) DeleteCatPicController(ctx *gin.Context) {

	catPicID := ctx.Params.ByName("cat_pic_id")
	if catPicID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           "Failed to get Image by this ImageID",
			"internal_error ": "Cat Image ID is not provided",
		})
		return
	}

	statusCode, result, err := cpc.CatPicHubService.DeleteCatPicService(catPicID)
	if err != nil {
		log.Println("Error while updating image in database : ", err)
		ctx.AbortWithStatusJSON(statusCode, result)
	} else {
		ctx.JSON(statusCode, result)
	}
}
