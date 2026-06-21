package controllers

import (
	"fmt"
	"gin-app/internal/dto"
	"gin-app/internal/repository"
	"gin-app/internal/services"
	"gin-app/internal/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControllers struct {
	s         *services.AuthService
	Validator validator.CustomValidator
	R         *repository.TokenRepository
}

func NewAuthControllers(s *services.AuthService, validator validator.CustomValidator, r *repository.TokenRepository) *AuthControllers {
	return &AuthControllers{
		s:         s,
		Validator: validator,
		R:         r,
	}
}

func (h *AuthControllers) LoginAdmin(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := h.Validator.Validate(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		1*24*60*60, 
		"/",
		"",
		true,
		true, 
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
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token not found"})
		return
	}

	err = h.R.DeleteToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete refresh token"})
		return
	}

	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

func (h *AuthControllers) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "refresh token tidak ditemukan",
		})
		return
	}

	result, err := h.s.RefreshToken(refreshToken)
	if err != nil {
		fmt.Println("REFRESH TOKEN ERROR :", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": result,
	})
}


