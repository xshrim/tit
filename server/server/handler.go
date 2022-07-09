package server

import (
	"context"
	"fmt"
	"net/http"
	"titserver/database"
	"titserver/model"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	user := c.MustGet("user").(model.User)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{"user": user},
	})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Test(c *gin.Context) {
	mgo := database.New("tit", "project")
	count, _ := mgo.CollectionCount()
	c.String(http.StatusOK, fmt.Sprintf("%d", count))
}

func UserInfo(c *gin.Context) {
	var user model.User
	json := make(map[string]any)
	_ = c.BindJSON(&json)
	result := database.New("tit", "user").FindOne("id", json["id"])
	_ = result.Decode(&user)
	fmt.Println(user)
}

func UserList(c *gin.Context) {
	//var users []model.User
	json := make(map[string]any)
	_ = c.BindJSON(&json)

	result, _ := database.New("tit", "user").FindMany(json)
	var users []model.User
	result.All(context.TODO(), &users)
	fmt.Println(users)

}

func Issue(c *gin.Context) {
	var uq model.UserReq
	if err := c.ShouldBind(&uq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 601,
			"msg":  "invalid parameters",
		})
	}

	var user model.User

	mgo := database.New("tit", "user")
	if uq.ID != "" {

		_ = mgo.FindOne("id", uq.ID).Decode(&user)
	} else if uq.Email != "" {
		_ = mgo.FindOne("email", uq.Email).Decode(&user)
	}

	if uq.Passwd == user.Passwd {
		tokenString, _ := GenToken(user)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  "authenticate failed",
	})
}
