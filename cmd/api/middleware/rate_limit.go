package middleware

import (
    "time"

    ginMerry "github.com/utrack/gin-merry"
    limit "github.com/yangxikun/gin-limit-by-key"

    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
)

// 针对IP每秒限速N
func RateLimit(N int) gin.HandlerFunc {
    return func(c *gin.Context) {
        limit.NewRateLimiter(func(c *gin.Context) string {
            return c.ClientIP() // limit rate by client ip
        }, func(c *gin.Context) (*rate.Limiter, time.Duration) {
            return rate.NewLimiter(rate.Every(1000*time.Millisecond), N), time.Hour // limit 10 qps/clientIp and permit bursts of at most 10 tokens, and the limiter liveness time duration is 1 hour
        }, func(c *gin.Context) {
            c.AbortWithStatus(429) // handle exceed rate limit request
        })
    }
}

// 异常打印
func Print() gin.HandlerFunc {
    return ginMerry.New(true).Handler()
}
