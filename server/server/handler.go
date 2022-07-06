package server

import "net/http"
import "github.com/gin-gonic/gin"

var handle map[string]func(*gin.Context)

func init() {
	handle = make(map[string]func(*gin.Context))
	handle["index"] = func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	}

	handle["ping"] = func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}
