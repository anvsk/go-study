package middleware

import (
    "encoding/json"
    "go-ticket/cmd/api/common"
    "go-ticket/pkg/store/db"
    "time"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
    "github.com/pieterclaerhout/go-log"
)

type login struct {
    Username string `form:"username" json:"username" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

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
        Authenticator: func(c *gin.Context) (interface{}, error) {
            var loginVals login
            if err := c.ShouldBind(&loginVals); err != nil {
                return "", jwt.ErrMissingLoginValues
            }
            user := common.User{}
            db.Orm.Raw("select * from users where username=? and password=? limit 1", loginVals.Username, loginVals.Password).Scan(&user)
            if user.ID > 0 {
                return user, nil
            }
            return nil, jwt.ErrFailedAuthentication
        },
        // login后组织payload
        // 这里往payload里面加信息，用户信息和其他信息
        PayloadFunc: func(data interface{}) jwt.MapClaims {
            payload := jwt.MapClaims{}
            if v, ok := data.(common.User); ok {
                tmp, _ := json.Marshal(common.JwtPayload{User: v})
                json.Unmarshal(tmp, &payload)
                return payload
            }
            return jwt.MapClaims{}
        },
        // 解析出payload
        IdentityHandler: func(c *gin.Context) interface{} {
            claims := jwt.ExtractClaims(c)
            return claims
        },
        // 权限校验 data=IdentityHandler返回值
        Authorizator: func(data interface{}, c *gin.Context) bool {
            v, ok := data.(jwt.MapClaims)
            if ok && v["Username"] == "admin" {
                return true
            }
            return false
        },
        // 权限不通过
        Unauthorized: func(c *gin.Context, code int, message string) {
            c.JSON(code, gin.H{
                "code":    405,
                "message": "Unauthorized",
            })
        },
        // TokenLookup is a string in the form of "<source>:<name>" that is used
        // to extract token from the request.
        // Optional. Default value "header:Authorization".
        // Possible values:
        // - "header:<name>"
        // - "query:<name>"
        // - "cookie:<name>"
        // - "param:<name>"
        TokenLookup: "header: Authorization, query: token, cookie: jwt",
        // TokenLookup: "query:token",
        // TokenLookup: "cookie:token",

        // TokenHeadName is a string in the header. Default value is "Bearer"
        TokenHeadName: "Bearer",

        // TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
        TimeFunc: time.Now,
    })

    if err != nil {
        log.Fatal("JWT Error:" + err.Error())
    }

    // When you use jwt.New(), the function is already automatically called for checking,
    // which means you don't need to call it again.
    errInit := Authjwt.MiddlewareInit()

    if errInit != nil {
        log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
    }
    return Authjwt
}
