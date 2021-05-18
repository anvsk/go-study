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
        PayloadFunc: func(data interface{}) jwt.MapClaims {
            log.InfoDump(111, "PayloadFunc")
            log.InfoDump(data, "data")
            payload := jwt.MapClaims{}
            if v, ok := data.(*common.User); ok {
                tmp, _ := json.Marshal(v)
                json.Unmarshal(tmp, &payload)
                return payload
            }
            return jwt.MapClaims{}
        },
        // IdentityHandler: func(c *gin.Context) interface{} {
        //     claims := jwt.ExtractClaims(c)
        //     log.InfoDump(111, "IdentityHandler")
        //     return common.User(claims)
        //     // return &common.User{
        //     //     UserName: claims[common.IdentityKey].(string),
        //     // }
        // },
        // LoginHandler 登录验证方法
        Authenticator: func(c *gin.Context) (interface{}, error) {
            log.InfoDump(111, "Authenticator")

            var loginVals login
            if err := c.ShouldBind(&loginVals); err != nil {
                return "", jwt.ErrMissingLoginValues
            }
            username := loginVals.Username
            password := loginVals.Password
            user := common.User{}
            db.Orm.Raw("select * from users where username=? and password=? limit 1", username, password).Scan(&user)
            log.DebugDump(user, "select user")
            if user.ID > 0 {
                return user, nil
            }
            return nil, jwt.ErrFailedAuthentication
        },
        // 权限校验
        Authorizator: func(data interface{}, c *gin.Context) bool {
            log.InfoDump(111, "Authorizator")
            if v, ok := data.(*common.User); ok && v.Username == "admin" {
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
