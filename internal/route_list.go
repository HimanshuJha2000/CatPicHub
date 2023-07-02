package internal

import (
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	// Setting up username and password from environment
	username := os.Getenv("CATPICHUB_USERNAME")
	password := os.Getenv("CATPICHUB_PASSWORD")

	// Middleware for basic authentication
	r.Use(func(ctx *gin.Context) {
		gin.BasicAuth(gin.Accounts{
			username: password,
		})(ctx)
		// Check if the response status is 401 Unauthorized
		if ctx.Writer.Status() == http.StatusUnauthorized {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Wrong username and password is passed!",
			})
			return
		}
	})

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
