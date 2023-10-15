package router

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	v1 "go-boss/api/http/v1"
	"net/http"
)

func InitApiRouter(engine *gin.Engine) {
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ginI18n.MustGetMessage(ctx, "welcome"))
	})

	versionOne := engine.Group("/v1")
	{
		versionOne.GET("/ping", v1.Ping)
		versionOne.GET("/admin", v1.GetAdmin)
	}

}
