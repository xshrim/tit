package server

import "github.com/gin-gonic/gin"

func route(r *gin.Engine) {

	r.GET("/", Index)

	v1 := r.Group("/api/v1")

	v1.POST("/login", Login)
	v1.POST("/auth", Auth)

	// {
	// 	auth := v1.Group("/auth")
	// }

	{
		user := v1.Group("/user")
		user.POST("/info", JWTAuthMiddleWare(), UserInfo)
		user.POST("/list", JWTAuthMiddleWare(), UserList)
	}

	{
		project := v1.Group("/project")
		project.POST("/info", JWTAuthMiddleWare(), ProjectInfo)
		project.POST("/list", JWTAuthMiddleWare(), ProjectList)
	}

	{
		rvitem := v1.Group("/rvitem")
		rvitem.POST("/info", JWTAuthMiddleWare(), RvitemInfo)
		rvitem.POST("/list", JWTAuthMiddleWare(), RvitemList)
	}

	{
		misc := v1.Group("/misc")
		misc.GET("/ping", Ping)
		misc.GET("/test", Test)
	}
}
