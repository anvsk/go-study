package middleware

import (
    "go-ticket/cmd/api/common"
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
        IdentityKey: common.IdentityKey,
        // LoginHandler 登录验证方法
        Authenticator: common.Authenticator,
        // login后组织payload
        // 这里往payload里面加信息，用户信息和其他信息
        PayloadFunc: common.PayloadFunc,

        /**********         校验token       ***********/

        // 权限校验 data=IdentityHandler返回值
        Authorizator: func(data interface{}, c *gin.Context) bool {
            if data.(float64) < 1 {
                return false
            }
            return true
        },
        // 权限不通过
        // Unauthorized: func(c *gin.Context, code int, message string) {
        //     c.JSON(code, gin.H{
        //         "code":    405,
        //         "message": "Unauthorized",
        //     })
        // },
    })

    if err != nil {
        log.Fatal("JWT Init Error:" + err.Error())
    }

    return Authjwt
}
