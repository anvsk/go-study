package router

import (
	"go-study/cmd/api/internal/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouters(e *gin.Engine) {
	e.GET("/users", admin.UsersHandler)
}
