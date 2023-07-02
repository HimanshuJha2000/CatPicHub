package controller

import (
	"fmt"
	"github.com/DevtronLabs/CatPicHub/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type catPicHubController struct {
	catPicHubService service.Service
}

var CatPicHub catPicHubController

func (cpc *catPicHubController) GetCatPicController(ctx *gin.Context) {

	value, err := fmt.Println("hello world")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, value)
	} else {
		ctx.JSON(http.StatusOK, value)
	}
}
