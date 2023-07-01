package internal

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {

	r := gin.Default()
	grp := r.Group("/api/")
	{
		grp.GET("cat", controller.Players.FetchPlayerList)
		grp.GET("cat/:cat_id", controller.Players.FetchPlayerByID)
		grp.GET("cat/paginate/:page_no", controller.Players.FetchPlayerByPage)
		grp.POST("cat", controller.Players.SleepService)
	}
	return r
}
