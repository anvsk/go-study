package middleware

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

type Context struct {
    Ctx *gin.Context
}

type response struct {
    Code    int         `json:"code"`
    Success bool        `json:"success"`
    Content interface{} `json:"content"`
    Message interface{} `json:"msg"`
}

func (c *Context) Response(code int, msg interface{}, content interface{}) {
    if msg == nil {
        c.Ctx.JSON(http.StatusOK, response{
            Code:    code,
            Success: true,
            Content: content,
            Message: msg,
        })
        return
    }
    // 错误格式
    c.Ctx.JSON(http.StatusOK, response{
        Code:    code,
        Success: false,
        Content: content,
        Message: msg,
    })
}

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
