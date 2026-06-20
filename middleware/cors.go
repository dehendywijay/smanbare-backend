package middleware

import (
	"fmt"
	"net/http"


	"github.com/gin-gonic/gin"
)



func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CORS middleware hit:", c.Request.Method, c.Request.URL.Path)

		c.Header("Access-Control-Allow-Origin", "https://web-sma-xi.vercel.app")
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

