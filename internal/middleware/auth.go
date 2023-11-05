package middleware

import (
	"github.com/gin-gonic/gin"
	"go-boss/pkg/auth/jwt"
	"log"
	"net/http"
	"strings"
)

func Authenticate() func(c *gin.Context) {
	return func(c *gin.Context) {
		//获取到请求头中的token
		authHeader := c.Request.Header.Get("Authorization")
		log.Println(authHeader)
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "访问失败,请登录!",
				"data": nil,
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "访问失败,无效的token,请登录!",
				"data": nil,
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, &gin.H{
				"code": 200,
				"msg":  "访问失败,无效的token,请登录!",
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Set("username", mc.Username)
		c.Next()
	}
}
