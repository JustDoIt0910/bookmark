package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Context)  {
	c.JSON(http.StatusOK, "pong")
}
