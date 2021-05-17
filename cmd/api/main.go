package main

import (
    "fmt"
    "go-ticket/cmd/api/internal/admin"
    "go-ticket/cmd/api/internal/shop"
    "go-ticket/pkg/util"

    "github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

func main() {
    util.InitUtil()
    // cache.InitCache()
    // db.InitDB()
    initServer(shop.Routers, admin.Routers) // 注册功能模块路由
}

func initServer(opts ...Option) {
    r := gin.New()
    options := []Option{}
    options = append(options, opts...)
    for _, opt := range options {
        opt(r)
    }
    if err := r.Run(); err != nil {
        fmt.Println(err)
    }
}
