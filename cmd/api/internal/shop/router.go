package shop

import (
    "fmt"
    "go-ticket/cmd/api/middleware"
    "time"

    "github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
    middleware.InitJWT()
    e.POST("login", middleware.Authjwt.LoginHandler)

    g := e.Group("goods")
    {
        g.GET("/", goodsHandler)
    }

    p := e.Group("user")
    p.Use(middleware.Authjwt.MiddlewareFunc())
    {
        p.GET("info", infoHandler)

    }
}

// 定义中间
func myTime(c *gin.Context) {
    start := time.Now()
    c.Next()
    // 统计时间
    since := time.Since(start)
    fmt.Println("程序用时：", since)
}

func myTime2(c *gin.Context) {
    start := time.Now()
    c.Next()
    // 统计时间
    since := time.Since(start)
    fmt.Println("程序用时2：", since)
}
