package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"sync"
	"time"
)

const (
	TOKEN_SECRET_KEY   = "secret"         // 密钥
	TOKEN_EXPIRE_TIME  = 2 * time.Hour    // 2小时过期
	TOKEN_REFRESH_TIME = 10 * time.Minute // 接近过期时会在header里面加上新token，客户端可以识别也可以自行拉取新
)

var Tb *TokenBucket

func init() {
	Tb = NewTokenBucket(100, 5000)
}

/**
 * 登陆
 * curl 127.0.0.1:8080/login -X POST
 * return {"token":"eyJhbGciO..."}
 */
func LoginHandler(c *gin.Context) {
	// TODO 验证账户密码
	// account + passwd 需要从db中拉取信息校验
	// 获取用户信息
	userId := "123"
	userName := "test 123"

	// 签名JWT
	tokenString, err := generateJWTToken(userId, userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	// 返回JWT给客户端
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

/**
 * 用户信息
 * curl 127.0.0.1:8080/v1/user -H "token:eyJhbGciO..."
 * return {"UserId":"123","UserName":"test 123","exp":1694741333,"nbf":1694734133,"iat":1694734133}
 */
func UserHandler(c *gin.Context) {
	claims, bool := c.Get("claims")
	if bool {
		// TODO 其他用户信息可以用UID查 缓存 和 数据库
		// findbyId()
		c.JSON(http.StatusOK, claims)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "not found"})
}

/**
 * 刷新token
 * curl 127.0.0.1:8080/v1/refresh-token -H "token:eyJhbGciO..."
 * return {"token":"eyJhbGciO..."}
 */
func RefreshTokenHandler(c *gin.Context) {
	claims, bool := c.Get("claims")
	if !bool {
		c.JSON(http.StatusOK, gin.H{"message": "not found claims"})
		return
	}
	fmt.Println(claims)
	val, ok := claims.(*jwtClaims)

	if !ok {
		c.JSON(http.StatusOK, gin.H{"message": "not found"})
		return
	}
	tokenString, err := generateJWTToken(val.UserId, val.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

// jwt
type jwtClaims struct {
	UserId               string
	UserName             string
	jwt.RegisteredClaims // jwt中标准格式
}

/**
 * 校验token
 * 如果想从服务端控制发出的token，可以通过redis标记也能达到让指定token提前过期的目的
 */
func AuthReqMiddWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取TOKEN
		tokenStr := c.GetHeader("token")
		if tokenStr == "" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Token not exist"})
			c.Abort()
			return
		}
		// 解析token
		token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(TOKEN_SECRET_KEY), nil
		})
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*jwtClaims)
		// 这里默认会检查ExpiresAt是否过期
		if ok && token.Valid {
			now := time.Now()
			// 检查过期时间，对快要过期的添加http header `refresh-token`
			if t := claims.ExpiresAt.Time.Add(-TOKEN_REFRESH_TIME); t.Before(now) {
				tokenString, err := generateJWTToken(claims.UserId, claims.UserName)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
					c.Abort()
					return
				}
				c.Header("refresh-token", tokenString) //
			}
			c.Set("claims", claims)
		}
	}
}

// 生成JWT token
func generateJWTToken(userId, userName string) (string, error) {
	now := time.Now()
	claims := jwtClaims{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(TOKEN_EXPIRE_TIME)}, // 过期时间
			IssuedAt:  jwt.NewNumericDate(now),                                   // 签发时间
			NotBefore: jwt.NewNumericDate(now),                                   // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(TOKEN_SECRET_KEY))
}

// 限流
func RateMiddWare(tb *TokenBucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !tb.AllowRequest() {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": http.StatusText(http.StatusTooManyRequests)})
			c.Abort()
			return
		}

	}
}

// 令牌
type TokenBucket struct {
	cap      int       // 桶容量
	rate     float64   // 每秒生产个数
	tokenNum int       // 当前计数
	lastTime time.Time // 上一个产生时间
	mu       sync.Mutex
}

func NewTokenBucket(cap int, rate float64) *TokenBucket {
	return &TokenBucket{
		cap:      cap,
		rate:     rate,
		tokenNum: cap,
		lastTime: time.Now(),
	}
}

// 拿令牌
func (tb *TokenBucket) AllowRequest() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	second := now.Sub(tb.lastTime).Seconds() // 计算经过多少秒
	newTokens := int(second * tb.rate)       // 计算产生的令牌数量

	if newTokens > 0 {
		tb.tokenNum = tb.tokenNum + newTokens
		if tb.tokenNum > tb.cap { // 不能超过容量
			tb.tokenNum = tb.cap
		}
		tb.lastTime = now
	}

	if tb.tokenNum > 0 {
		tb.tokenNum--
		return true
	}

	return false
}
