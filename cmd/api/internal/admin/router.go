package admin

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
    e.GET("/users", usersHandler)
}
