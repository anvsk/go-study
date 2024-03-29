package router

import (
	"go-study/cmd/api/internal/shop"
	"go-study/cmd/api/middleware"

	"github.com/gin-gonic/gin"
)

func ShopRouters(e *gin.Engine) {
	e.Use(middleware.Cors())
	e.POST("wechat/official-account/menu", shop.TestMenu)

	middleware.InitJWT()
	e.POST("login", middleware.Authjwt.LoginHandler)

	g := e.Group("goods")
	{
		g.GET("/", shop.GoodsHandler)
	}

	p := e.Group("user")
	p.Use(middleware.Authjwt.MiddlewareFunc())
	{
		p.GET("info", shop.InfoHandler)

	}
}
