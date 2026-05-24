package middleware

import (
	"fmt"
	"net/http"
	"os"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func getAccessSecret() []byte {
	secret := os.Getenv("ACCESS_SECRET_KEY")
	if secret == "" {
		return []byte("ACCESS_SECRET_KEY")
	}
	return []byte(secret)
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CORS middleware hit:", c.Request.Method, c.Request.URL.Path)

		c.Header("Access-Control-Allow-Origin", "https://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header kosong"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return getAccessSecret(), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "access token tidak valid"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["type"] != "access" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token bukan access token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
