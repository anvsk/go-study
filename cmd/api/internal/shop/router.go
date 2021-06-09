package shop

import (
    "go-ticket/cmd/api/middleware"

    "github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
    e.Use(middleware.Cors())
    e.GET("wechat/official-account/menu", testMenu)

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
