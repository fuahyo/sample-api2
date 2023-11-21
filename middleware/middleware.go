package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

type MiddlewareAuth struct {
	DB *sqlx.DB
}

const (
	SECRET = "secret"
)

func (h *MiddlewareAuth) AuthValidation(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token required",
		})
		c.Abort()
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, invalid := token.Method.(*jwt.SigningMethodHMAC); !invalid {
			return nil, fmt.Errorf("invalid token", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "you are not authorized",
			"error":   err.Error(),
		})
	}
}
