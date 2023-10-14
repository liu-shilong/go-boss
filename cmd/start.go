package cmd

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-boss/internal/model"
	"go-boss/router"
	"log"
	"time"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "启动服务",
	Example: "./go-boss start",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap()
	},
}

func init() {

}

// 启动
func bootstrap() {
	// 初始化数据库
	model.InitDB()
	// 启动Gin引擎
	engine()
}

func engine() {

	engine := gin.Default()
	// 跨域处理
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// 初始化路由
	router.InitApiRouter(engine)

	// 运行
	err := engine.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
