package cache

import (
    "time"

    "github.com/patrickmn/go-cache"
)

type GoCache struct {
    cache *cache.Cache
}

func NewGoCache() CacheInterface {
    return &GoCache{
        cache: cache.New(5*time.Minute, 10*time.Minute),
    }
}

func (c *GoCache) Set(key string, value interface{}, d time.Duration) error {
    c.cache.Set(key, value, d)
    return nil
}

func (c *GoCache) Get(key string) interface{} {
    if x, found := c.cache.Get(key); found {
        return x
    }
    return nil
}
