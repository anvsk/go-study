package main

import (
    "go-ticket/cmd/api/common"
    "go-ticket/cmd/api/internal/admin"
    "go-ticket/cmd/api/internal/shop"
    "go-ticket/pkg/store/cache"
    "go-ticket/pkg/store/db"
    "go-ticket/pkg/util"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "golang.org/x/sync/errgroup"
)

type Option func(*gin.Engine)

var (
    g errgroup.Group
)

func newRouter(opts ...Option) http.Handler {
    e := gin.New()
    e.Use(gin.Logger(), gin.Recovery())
    for _, opt := range opts {
        opt(e)
    }
    return e
}

// 多服务端口
func main() {
    // 加载配置文件
    util.InitUtil()
    // 连接数据库
    db.InitDB()
    // 连接数据库
    cache.InitCache()
    // 生成测试用户表
    common.InitUserTable()

    server01 := &http.Server{
        Addr:         ":8081",
        Handler:      newRouter(admin.Routers),
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    server02 := &http.Server{
        Addr:         ":8080",
        Handler:      newRouter(shop.Routers),
        ReadTimeout:  2 * time.Second,
        WriteTimeout: 2 * time.Second,
    }
    // f, _ := os.Create("gin1.log")
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
    // gin.ForceConsoleColor()
    g.Go(func() error {
        return server01.ListenAndServe()
    })

    g.Go(func() error {
        return server02.ListenAndServe()
    })

    if err := g.Wait(); err != nil {
        log.Fatal(err)
    }

}
