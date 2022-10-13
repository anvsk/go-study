package cache

import (
	"fmt"
	"go-study/pkg/util"

	"time"

	"github.com/go-redis/redis"
	"github.com/pieterclaerhout/go-log"
)

type RedisCache struct {
	cache *redis.Client
}

func NewRedisCache() CacheInterface {
	cf := util.Config.Store.Cache.Connects["redis"]
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		Password: cf.Password, // no password set
		DB:       cf.DbIndex,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.ErrorDump(err, "redis ping error")
	}
	return &RedisCache{
		cache: client,
	}
}

func (c *RedisCache) Set(key string, value interface{}, d time.Duration) error {
	return c.cache.Set(key, value, d).Err()
}

func (c *RedisCache) Get(key string) interface{} {
	x, err := c.cache.Get(key).Result()
	if err != nil {
		return nil
	}
	return x
}

type Redis struct {
	Addr string
	Pass string
}
