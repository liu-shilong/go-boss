package cmd

import (
	"encoding/json"
	"github.com/gin-contrib/cors"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-boss/internal/model"
	"go-boss/pkg/cache"
	"go-boss/pkg/database/mongo"
	"go-boss/router"
	"golang.org/x/text/language"
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

// 启动
func bootstrap() {
	// 初始化数据库
	model.Connect()
	cache.InitRedis()
	mongo.InitMongo()
	// 启动Gin引擎
	engine()
}

func engine() {
	// 运行模式 debug
	gin.SetMode(gin.DebugMode)
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

	// 多语言
	engine.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		RootPath:         "./localize",
		AcceptLanguage:   []language.Tag{language.SimplifiedChinese, language.English},
		DefaultLanguage:  language.SimplifiedChinese,
		UnmarshalFunc:    json.Unmarshal,
		FormatBundleFile: "json",
	})))

	// 初始化路由
	router.InitApiRouter(engine)

	// 运行
	err := engine.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
