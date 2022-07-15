package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"tit/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const jwtSecret string = "xshrim"
const tokenExpireDuration = time.Hour * 3

type claim struct {
	Uid   string `json:"uid"`
	Name  string `json:"name"`
	Uname string `json:"uname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Dept  string `json:"dept"`
	Group string `json:"group"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// func Secret() jwt.Keyfunc {
// 	return func(token *jwt.Token) (interface{}, error) {
// 		return []byte("手写的从前"), nil // 这是我的secret
// 	}
// }

func GenToken(user model.User) (string, error) {
	c := claim{user.Uid, user.Name, user.Uname, user.Email, user.Phone, user.Dept, user.Group, user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)),
			Issuer:    "tit",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(jwtSecret))
}

func ParseToken(tokenString string) (*claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claim{}, func(token *jwt.Token) (i any, err error) {
		return []byte(jwtSecret), nil
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
			Uid:   claim.Uid,
			Name:  claim.Name,
			Uname: claim.Uname,
			Email: claim.Email,
			Phone: claim.Phone,
			Dept:  claim.Dept,
			Group: claim.Group,
			Role:  claim.Role,
		})
		c.Next()
	}
}
