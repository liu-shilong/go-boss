package service

import (
	"github.com/gin-gonic/gin"
	"go-boss/internal/model"
	"go-boss/pkg/auth/jwt"
	"go-boss/pkg/http/response"
)

func GetAdmin(c *gin.Context) {
	admin := model.FindAdminById()
	data := make(map[string]interface{})
	data["name"] = admin.Name
	data["mobile"] = admin.Mobile
	response.Success(c, data)
}

func Login(c *gin.Context) {

	username := "admin"
	password := "123456"

	tokenString, err := jwt.GenerateToken(username, password)
	if err != nil {
		response.Fail(c)
		return
	}

	data := make(map[string]interface{})
	data["token"] = tokenString
	response.Success(c, data)
}
