package internal

import (
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/controller"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	username := os.Getenv("CATPICHUB_USERNAME")
	password := os.Getenv("CATPICHUB_PASSWORD")

	// Middleware for basic authentication
	r.Use(gin.BasicAuth(gin.Accounts{
		username: password,
	}))

	grp := r.Group("/api/")
	{
		grp.POST("cat-pics", controller.CatPicHub.CreateCatPicController)
		grp.GET("cat-pics/:cat_pic_id", controller.CatPicHub.GetCatPicByIDController)
		grp.GET("cat-pics/list/:page_no/:page_size", controller.CatPicHub.GetCatPicListController)
		grp.PUT("cat-pics/:cat_pic_id", controller.CatPicHub.UpdateCatPicController)
		grp.DELETE("cat-pics/:cat_pic_id", controller.CatPicHub.DeleteCatPicController)
	}
	return r
}
