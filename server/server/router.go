package server

import "github.com/gin-gonic/gin"

func route(r *gin.Engine) {

	r.GET("/", JWTAuthMiddleWare(), Index)

	v1 := r.Group("/api/v1")

	// {
	// 	auth := v1.Group("/auth")
	// }

	{
		user := v1.Group("/user")
		user.POST("/info", UserInfo)
		user.POST("/list", UserList)
	}

	{
		misc := v1.Group("/misc")
		misc.GET("/ping", Ping)
		misc.GET("/test", Test)
	}
}
