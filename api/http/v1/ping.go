package v1

import (
	. "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-boss/pkg/cache"
	"go-boss/pkg/database/mongo"
	"go-boss/pkg/file/excel"
	"go-boss/pkg/mailer"
	"go-boss/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func Ping(c *gin.Context) {
	//testRedis()
	// testMongo()
	// testGeo()
	excel.Create()
	testEmail()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "pong",
	})
}

func testGeo() {
	geo := util.Ip2region("116.77.75.255")
	log.Printf("geo: %v \n", geo) // geo: map[city:深圳市 country:中国 isp:天威 province:广东省]
}

// 测试mongodb
func testMongo() {
	ctx := Background()
	client := mongo.MongoClient
	collection := client.Database("testing").Collection("numbers")

	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID

	log.Printf("objectId: %v \n", id)

}

// 测试redis
func testRedis() {
	ctx := Background()
	client := cache.RedisCache
	key := "go redis"
	err := client.Set(ctx, key, "test go redis", 0).Err()
	if err != nil {
		fmt.Print(err)
		return
	}
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func testEmail() {
	to := "1932255901@qq.com"
	subject := "test gmail"
	body := `<p style="color:red;">This is a email with gmail!</p>`
	// 发送邮件
	err := mailer.SendText(to, subject, body)
	if err != nil {
		println(err.Error())
	}
}
