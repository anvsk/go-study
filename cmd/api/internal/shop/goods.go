package shop

import (
    "go-ticket/cmd/api/com"
    "go-ticket/service/shop"

    "github.com/gin-gonic/gin"
)

func GoodsHandler(c *gin.Context) {
    req := shop.ReqList{}
    if err := c.Bind(&req); err != nil {
        com.Error(c, err.Error())
        return
    }
    com.Responce(c, shop.List(req))
}

func InfoHandler(c *gin.Context) {
    com.Responce(c, com.Uinfo(c))
}
