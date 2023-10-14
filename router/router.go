package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-boss/api/http/v1"
)

func InitApiRouter(engine *gin.Engine) {

	versionOne := engine.Group("/v1")
	{
		versionOne.GET("/ping", v1.Ping)
		versionOne.GET("/admin", v1.GetAdmin)
	}

}
