package v1

import (
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"go-boss/internal/service"
)

func GetAdmin(c *gin.Context) {

	service.GetAdmin(c)
}

func Login(c *gin.Context) {
	service.Login(c)
}

func resJson() string {
	js := simplejson.New()
	js.Set("name", "gosimplejson")
	js.Set("author", "go语言")
	js.Set("github", "https://github.com/bitly/go-simplejson")
	js.Set("stars", 3236)
	out, err := js.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}
	return string(out)
}
