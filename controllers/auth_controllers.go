package controllers

import (
	"gin-app/models"
	"gin-app/services"
	"gin-app/utility"
	"net/http"
	"os"
	

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)


func getRefreshSecret() []byte {
	secret := os.Getenv("REFRESH_SECRET_KEY")
	if secret == "" {
		return []byte("REFRESH_SECRET_KEY")
	}
	return []byte(secret)
}

func LoginAdmin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// var adminInput models.Admin

	// if err := c.ShouldBindJSON(&adminInput); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	// 	return
	// }


	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	admin, err := services.LoginAdmin(username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	c.SetCookie(
		"refresh_token",
		admin.RefreshToken,
		7*24*60*60, // 7 hari
		"/",
		"",
		false, // true kalau pakai HTTPS
		true,  // HttpOnly
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data": map[string]interface{}{
			"access_token": admin.AccessToken,
			"username":     admin.Username,
		},
	})
}

func CreateAdmin(c *gin.Context) {
	var adminInput models.Admin

	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	username := adminInput.Username
	password := adminInput.Password

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	admin, err := services.CreateAdmin(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin berhasil dibuat",
		"data":    admin,
	})
}

func LogoutAdmin(c *gin.Context) {
	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

func RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token tidak ditemukan"})
		return
	}

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return getRefreshSecret(), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token tidak valid"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != "refresh" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token bukan refresh token"})
		return
	}

	userIDFloat := claims["user_id"].(float64)
	userID := uint(userIDFloat)

	newAccessToken, err := utility.GenerateAccessToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat access token baru"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}