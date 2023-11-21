package auth

import (
	"log"
	"net/http"
	"sample-api2/models"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	DB *sqlx.DB
}

const (
	USER     = "admin"
	PASSWORD = "12345678"
	SECRET   = "secret"
)

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var user models.Credential
	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("Error fetching credentials from the request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch credentials"})
		return
	}

	if user.Username == USER && user.Password == PASSWORD {
		// Authentication successful
		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "test",
			IssuedAt:  time.Now().Unix(),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token, err := sign.SignedString([]byte(SECRET))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"token":   token,
		})
	} else {
		// Authentication failed
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user or password invalid",
		})
	}
}

// package auth

// import (
// 	"log"
// 	"net/http"
// 	"sample-api2/models"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	jwt "github.com/golang-jwt/jwt/v4"
// 	"github.com/jmoiron/sqlx"
// )

// type UserHandler struct {
// 	DB *sqlx.DB
// }

// const (
// 	USER     = "admin"
// 	PASSWORD = "galang"
// 	SECRET   = "secret"
// )

// func (h *UserHandler) LoginHandler(c *gin.Context) {
// 	var user models.Credential
// 	err := c.Bind(&user)
// 	if err != nil {
// 		log.Printf("Error binding credentials: %v", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	if user.Username == USER && user.Password == PASSWORD {
// 		// Authentication successful
// 		claim := jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
// 			Issuer:    "test",
// 			IssuedAt:  time.Now().Unix(),
// 		}

// 		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
// 		token, err := sign.SignedString([]byte(SECRET))
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			c.Abort()
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "success",
// 			"token":   token,
// 		})
// 	} else {
// 		// Authentication failed
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "user or password invalid",
// 		})
// 	}
// }
