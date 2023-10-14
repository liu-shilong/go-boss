package v1

import (
	"github.com/gin-gonic/gin"
	"go-boss/internal/service"
)

func GetAdmin(c *gin.Context) {
	service.GetAdminInfo(c)
}
