package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/gin-contrib/cors"
	ginI18n "github.com/gin-contrib/i18n"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-boss/internal/model"
	"go-boss/pkg/cache"
	"go-boss/pkg/database/mongo"
	"go-boss/router"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/text/language"
	"io"
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

	// 日志
	logger, _ := zap.NewProduction()
	engine.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		Context: ginzap.Fn(func(c *gin.Context) []zapcore.Field {
			fields := []zapcore.Field{}
			// log request ID
			if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
				fields = append(fields, zap.String("request_id", requestID))
			}

			// log trace and span ID
			if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
				fields = append(fields, zap.String("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()))
				fields = append(fields, zap.String("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()))
			}

			// log request body
			var body []byte
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)
			fields = append(fields, zap.String("body", string(body)))

			return fields
		}),
	}))

	// 性能分析
	// pprof.Register(engine)

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
