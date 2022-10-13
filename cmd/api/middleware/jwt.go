package middleware

import (
	"go-study/cmd/api/com"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/pieterclaerhout/go-log"
)

// User demo

var Authjwt *jwt.GinJWTMiddleware

func InitJWT() *jwt.GinJWTMiddleware {
	var err error
	// the jwt middleware
	Authjwt, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("78^*^*&SJFHJSDHFLS^&%$^"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: com.IdentityKey,
		// LoginHandler 登录验证方法
		Authenticator: com.Authenticator,
		// login后组织payload
		// 这里往payload里面加信息，用户信息和其他信息
		PayloadFunc: com.PayloadFunc,

		/**********         校验token       ***********/

		// 权限校验 data=IdentityHandler返回值
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if data.(float64) < 1 {
				return false
			}
			return true
		},
	})

	if err != nil {
		log.Fatal("JWT Init Error:" + err.Error())
	}

	return Authjwt
}
