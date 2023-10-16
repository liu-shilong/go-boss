package router

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	v1 "go-boss/api/http/v1"
	"go-boss/pkg/auth/jwt"
	"net/http"
)

func InitApiRouter(engine *gin.Engine) {

	// ping测试
	engine.GET("/ping", v1.Ping)
	// 多语言测试
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ginI18n.MustGetMessage(ctx, "welcome"))
	})

	// 限流
	engine.Use(jwt.RateMiddWare(jwt.Tb))
	// 登录
	engine.GET("/login", jwt.LoginHandler)

	// v1版本
	versionOne := engine.Group("/v1").Use(jwt.AuthReqMiddWare())
	{
		// 获取用户信息
		versionOne.GET("/user", jwt.UserHandler)
		// 刷新token
		versionOne.GET("/refresh-token", jwt.RefreshTokenHandler)
		// 获取admin-测试
		versionOne.GET("/admin", v1.GetAdmin)
	}

}
