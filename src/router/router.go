package router

import (
	"go-localize/docs"
	v1 "go-localize/src/controller/v1"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/v1/testapi/get-string-by-int/:some_id", v1.GetStringByInt)
	router.GET("/v1/testapi/get-struct-array-by-string/:some_id", v1.GetStructArrayByString)
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api/v1")
	{
		eg := api.Group("/example")
		{
			eg.GET("/helloworld", v1.Helloworld)
		}
		api.GET("/me", v1.GetMe)
		api.POST("/me", v1.PostMe)
		api.PUT("/me", v1.PutMe)
		api.DELETE("/me", v1.DeleteMe)
	}
	return router
}
