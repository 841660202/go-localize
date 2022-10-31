package router

import (
	"go-localize/docs"
	v1 "go-localize/src/controller/v1"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	router := gin.Default()
	// router.GET("/v1/testapi/get-string-by-int/:some_id", v1.GetStringByInt)
	// router.GET("/v1/testapi/get-struct-array-by-string/:some_id", v1.GetStructArrayByString)
	docs.SwaggerInfo.BasePath = "/api/v1"
	api_v1 := router.Group("/api/v1")
	{
		eg := api_v1.Group("/example")
		{
			eg.GET("/helloworld", v1.GetHelloworld)
		}
		me := api_v1.Group("/me")
		{
			me.GET("", v1.GetMe)
			me.POST("", v1.PostMe)
			me.PUT("", v1.PutMe)
			me.DELETE("", v1.DeleteMe)
		}

		examples := api_v1.Group("/examples")
		{
			examples.GET("ping", v1.PingExample)
			examples.GET("calc", v1.CalcExample)
			examples.GET("groups/:group_id/accounts/:account_id", v1.PathParamsExample)
			examples.GET("header", v1.HeaderExample)
			examples.GET("securities", v1.SecuritiesExample)
			examples.GET("attribute", v1.AttributeExample)
		}

		accounts := api_v1.Group("/accounts")
		{
			accounts.GET(":id", v1.ShowAccount)
			accounts.GET("", v1.ListAccounts)
			accounts.POST("", v1.AddAccount)
			accounts.DELETE(":id", v1.DeleteAccount)
			accounts.PATCH(":id", v1.UpdateAccount)
			accounts.POST(":id/images", v1.UploadAccountImage)
		}
	}
	return router
}
