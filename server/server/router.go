package server

import "github.com/gin-gonic/gin"

func route(r *gin.Engine) {
	r.GET("/", handle("index"))
	r.GET("/ping", handle("ping"))
}
