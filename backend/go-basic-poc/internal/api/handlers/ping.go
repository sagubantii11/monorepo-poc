package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) { // Use gin.Context
	c.String(http.StatusOK, "pong")
}
