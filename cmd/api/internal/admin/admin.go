package admin

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func UsersHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "users": []string{"xiaoming", "xiaohong"},
    })
}
