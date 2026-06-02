package controllers

import (
	"gin-app/internal/dto"
	"gin-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthControllers struct {
	s *services.AuthService
}

func NewAuthControllers(s *services.AuthService) *AuthControllers {
	return &AuthControllers{
		s: s,
	}
}

func (h *AuthControllers) LoginAdmin(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}


	if loginRequest.Username == "" || loginRequest.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	admin, err := h.s.LoginAdmin(loginRequest.Username, loginRequest.Password)
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
		true, // true kalau pakai HTTPS
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


func (h *AuthControllers) LogoutAdmin(c *gin.Context) {
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

func (h *AuthControllers) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	
	result, err := h.s.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": result,
	})
}