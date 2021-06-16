package router

import (
    "go-ticket/cmd/api/internal/admin"

    "github.com/gin-gonic/gin"
)

func AdminRouters(e *gin.Engine) {
    e.GET("/users", admin.UsersHandler)
}
