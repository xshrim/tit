package server

import (
	"context"
	"fmt"
	"strings"

	"net/http"
	"tit/database"
	"tit/model"

	"github.com/gin-gonic/gin"
)

// func Index(c *gin.Context) {

// 	user := c.MustGet("user").(model.User)

// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 200,
// 		"msg":  "success",
// 		"data": gin.H{"user": user},
// 	})
// }

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Login(c *gin.Context) {
	var uq model.UserReq

	if err := c.ShouldBind(&uq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 601,
			"msg":  "invalid parameters",
		})
	}

	var user model.User

	mgo := database.New("tit", "user")
	if uq.Uid != "" {
		_ = mgo.FindOne("uid", uq.Uid).Decode(&user)
	} else if uq.Email != "" {
		_ = mgo.FindOne("email", uq.Email).Decode(&user)
	}

	if uq.Passwd == user.Passwd {
		tokenString, err := GenToken(user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "generate token failed: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"token": tokenString,
				"uid":   user.Uid,
				"name":  user.Name,
				"uname": user.Uname,
				"email": user.Email,
				"phone": user.Phone,
				"dept":  user.Dept,
				"group": user.Group,
				"role":  user.Role},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  "authenticate failed",
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
	result := database.New("tit", "user").FindOne("uid", json["uid"])
	_ = result.Decode(&user)
	fmt.Println(user)
}

func UserList(c *gin.Context) {
	//var users []model.User
	json := make(map[string]any)
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 601,
			"msg":  "invalid parameters: " + err.Error(),
		})
		return
	}

	result, _ := database.New("tit", "user").FindMany(json)
	var users []model.User
	result.All(context.TODO(), &users)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": users,
	})

}

func ProjectInfo(c *gin.Context) {
	var project model.Project
	json := make(map[string]any)
	_ = c.BindJSON(&json)
	result := database.New("tit", "project").FindOne("pid", json["pid"])
	_ = result.Decode(&project)
	fmt.Println(project)
}

func ProjectList(c *gin.Context) {

	json := make(map[string]any)
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 601,
			"msg":  "invalid parameters: " + err.Error(),
		})
		return
	}

	result, _ := database.New("tit", "project").FindMany(json)
	var projects []model.Project
	result.All(context.TODO(), &projects)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": projects,
	})
}

func RvitemInfo(c *gin.Context) {
	var rvitem model.Rvitem
	json := make(map[string]any)
	_ = c.BindJSON(&json)
	result := database.New("tit", "rvitem").FindOne("rid", json["rid"])
	_ = result.Decode(&rvitem)
	fmt.Println(rvitem)
}

func RvitemList(c *gin.Context) {

	json := make(map[string]any)
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 601,
			"msg":  "invalid parameters: " + err.Error(),
		})
		return
	}

	result, _ := database.New("tit", "rvitem").FindMany(json)
	var rvitems []model.Rvitem
	result.All(context.TODO(), &rvitems)
	for idx, _ := range rvitems {
		rvitems[idx].Tips = strings.ReplaceAll(rvitems[idx].Tips, "\n", " ")
		rvitems[idx].Require = strings.ReplaceAll(rvitems[idx].Require, "\n", " ")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": rvitems,
	})
}

func Auth(c *gin.Context) {
	var uq model.UserReq
	if err := c.ShouldBind(&uq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 601,
			"msg":  "invalid parameters: " + err.Error(),
		})
		return
	}

	claim, err := ParseToken(uq.Token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 501,
			"msg":  "parse token failed: " + err.Error(),
		})
		return
	}

	if claim.Uid != uq.Uid || claim.Email != uq.Email {
		c.JSON(http.StatusOK, gin.H{
			"code": 501,
			"msg":  "user id or email not matched",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
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
	if uq.Uid != "" {
		_ = mgo.FindOne("uid", uq.Uid).Decode(&user)
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
