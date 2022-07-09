package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"titserver/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const jwtSecret = "www"
const tokenExpireDuration = time.Hour * 3

type claim struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Uname    string `json:"uname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Dept     string `json:"dept"`
	Group    string `json:"group"`
	Reviewer string `json:"reviewer"`
	jwt.StandardClaims
}

func GenToken(user model.User) (string, error) {
	c := claim{user.ID, user.Name, user.Uname, user.Email, user.Phone, user.Dept, user.Group, user.Reviewer,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(),
			Issuer:    "tit",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claim{}, func(token *jwt.Token) (i any, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*claim); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func JWTAuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 602,
				"msg":  "authorization empty in http header",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claim, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("user", model.User{
			ID:       claim.ID,
			Name:     claim.Name,
			Uname:    claim.Uname,
			Email:    claim.Email,
			Phone:    claim.Phone,
			Dept:     claim.Dept,
			Group:    claim.Group,
			Reviewer: claim.Reviewer,
		})
		c.Next()
	}
}
