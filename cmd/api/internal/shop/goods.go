package shop

import (
    "go-ticket/service/shop"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func goodsHandler(c *gin.Context) {
    <-time.After(200 * time.Millisecond)
    req := shop.ReqList{}
    if err := c.Bind(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"params error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": shop.List(req),
    })
}

func infoHandler(c *gin.Context) {
    //
    c.JSON(http.StatusOK, gin.H{
        "data": shop.Info(),
    })
}
