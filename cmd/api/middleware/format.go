package middleware

import (
    "fmt"
    "time"

    "github.com/gin-gonic/gin"
)

func ConsoleLogFormat() gin.HandlerFunc {
    return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

        // your custom format
        return fmt.Sprintf("%s - [%s] %s %s %d %s  \n%s\n",
            param.ClientIP,
            param.TimeStamp.Format(time.Kitchen),
            param.Method,
            param.Path,
            // param.Request.Proto,
            param.StatusCode,
            param.Latency,
            // param.Request.UserAgent(),
            param.ErrorMessage,
        )
    })
}
