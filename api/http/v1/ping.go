package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-boss/pkg/cache"
	"net/http"
)

func Ping(c *gin.Context) {
	ctx := context.Background()
	client := cache.RedisCache
	err := client.Set(ctx, "111", "222", 0).Err()
	if err != nil {
		fmt.Print(err)
		return
	}
	val, err := client.Get(ctx, "111").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "pong",
	})
}
