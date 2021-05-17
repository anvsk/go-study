package shop

import (
    "fmt"
    "go-ticket/cmd/api/middleware"
    "time"

    "github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
    g := e.Group("goods")
    g.Use(myTime)
    g.Use(middleware.RateLimit(1))
    g.Use(middleware.Print())
    g.Use(myTime2)
    {
        g.GET("/", goodsHandler)
    }

    p := e.Group("user")
    {
        p.GET("/info", infoHandler)

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
