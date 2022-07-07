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
	c.String(http.StatusOK, "Hello World!")
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
