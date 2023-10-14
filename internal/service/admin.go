package service

import (
	"github.com/gin-gonic/gin"
	"go-boss/internal/model"
	"go-boss/pkg/http/response"
)

func GetAdminInfo(c *gin.Context) {
	admin := model.FindAdminById()
	data := make(map[string]interface{})
	data["name"] = admin.Name
	data["mobile"] = admin.Mobile
	response.Success(c, data)
}
