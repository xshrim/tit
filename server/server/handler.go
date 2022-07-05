package server

import "net/http"
import "github.com/gin-gonic/gin"

func handle(hd string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	}

	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}
