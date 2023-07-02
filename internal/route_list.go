package internal

import (
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	grp := r.Group("/api/")
	{
		grp.POST("cat-pics", controller.CatPicHub.CreateCatPicController)
		grp.GET("cat-pics/:cat_pic_id", controller.CatPicHub.GetCatPicByIDController)
		grp.GET("cat-pics/list", controller.CatPicHub.GetCatPicListController)
		grp.PUT("cat-pics/:cat_pic_id", controller.CatPicHub.UpdateCatPicController)
		grp.DELETE("cat-pics/:cat_pic_id", controller.CatPicHub.DeleteCatPicController)
	}
	return r
}
