package router

import (
	"github.com/gin-gonic/gin"
	"hatgo/pkg/conf"
	"hatgo/app/router/api/v1"
	"hatgo/app/middle"
)

func InitRouter() *gin.Engine {
	gin.SetMode(conf.RunMode)
	r := gin.New()
	r.Use(gin.Recovery())
	if conf.RunMode == gin.DebugMode {
		r.Use(gin.Logger())
	}
	r.Use(middle.Core, middle.TouchBody)

	r.POST("/login", v1.RLogin)

	ws := r.Group("/")
	{
		ws.GET("socket", v1.RWebSocket)
	}
	api := r.Group("/")
	{
		api.POST("get-test", v1.RGetTest)
		api.POST("add-test", v1.RAddTest)
		api.POST("upload", v1.RUpload)
		api.POST("get-xml", v1.GetXml)
	}
	return r
}
