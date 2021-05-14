package cache

import (
    "go-ticket/pkg/util"

    "time"
)

var C CacheInterface

type CacheInterface interface {
    Set(key string, value interface{}, d time.Duration) error
    Get(key string) interface{}
}

func InitCache() {
    driver := util.Config.Store.Cache.DefaultDriver
    switch driver {
    case "go-cache":
        C = NewGoCache()
    case "redis":
        C = NewRedisCache()
    }
}
