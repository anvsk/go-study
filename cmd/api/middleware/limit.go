package middleware

import (
    "time"

    llimit "github.com/aviddiviner/gin-limit"

    ginMerry "github.com/utrack/gin-merry"
    limit "github.com/yangxikun/gin-limit-by-key"

    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
)

// 针对IP限速
func RateLimitByKey(N int) gin.HandlerFunc {
    return func(c *gin.Context) {
        limit.NewRateLimiter(func(c *gin.Context) string {
            return c.ClientIP() // limit rate by client ip
        }, func(c *gin.Context) (*rate.Limiter, time.Duration) {
            // limit 10 qps/clientIp and permit bursts of at most 10 tokens,
            //  and the limiter liveness time duration is 1 hour
            return rate.NewLimiter(rate.Every(1*time.Millisecond), N), time.Hour
        }, func(c *gin.Context) {
            c.AbortWithStatus(429) // handle exceed rate limit request
        })
    }
}

// 限制同时处理的请求数
func RateLimitByMaxNum(N int) gin.HandlerFunc {
    return llimit.MaxAllowed(N)
}

// 同上
func RateLimitByChannel(N int) gin.HandlerFunc {
    sem := make(chan struct{}, N)
    acquire := func() { sem <- struct{}{} }
    release := func() { <-sem }
    return func(c *gin.Context) {
        acquire()       // before request
        defer release() // after request
        c.Next()
    }
}

// 异常打印
func Print() gin.HandlerFunc {
    return ginMerry.New(true).Handler()
}
