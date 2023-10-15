package v1

import (
	. "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-boss/pkg/cache"
	"go-boss/pkg/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func Ping(c *gin.Context) {
	//testRedis()
	testMongo()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "pong",
	})
}

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
