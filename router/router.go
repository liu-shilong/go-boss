package router

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	v1 "go-boss/api/http/v1"
	"go-boss/internal/middleware"
	"net/http"
)

func InitApiRouter(engine *gin.Engine) {

	// ping测试
	engine.GET("/ping", v1.Ping)
	// 多语言测试
	engine.GET("/", func(ctx *gin.Context) {
		ctx.Writer.Header().Add("X-Request-Id", "1234-5678-9012")
		ctx.String(http.StatusOK, ginI18n.MustGetMessage(ctx, "welcome"))
	})

	// 登录
	engine.GET("/login", v1.Login)

	// v1版本
	versionOne := engine.Group("/v1").Use(middleware.Authenticate())
	{

		// 获取admin-测试
		versionOne.GET("/admin", v1.GetAdmin)
	}

}
