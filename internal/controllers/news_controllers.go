package controllers

import (
	"gin-app/internal/dto"
	"gin-app/internal/services"
	"gin-app/internal/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsControllers struct {
	s         *services.NewsService
	Validator validator.CustomValidator
}

func NewNewsControllers(s *services.NewsService, validator validator.CustomValidator) *NewsControllers {
	return &NewsControllers{
		s:         s,
		Validator: validator,
	}
}

func (h *NewsControllers) CreateNews(c *gin.Context) {
	var news dto.NewsRequest
	if err := c.ShouldBind(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Validator.Validate(&news)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.s.CreateNews(news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berita berhasil dibuat",
	})
}

func (h *NewsControllers) GetNews(c *gin.Context) {
	result, err := h.s.GetNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *NewsControllers) GetNewsByID(c *gin.Context) {
	slug := c.Param("slug")

	result, err := h.s.GetNewsByID(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *NewsControllers) UpdateNews(c *gin.Context) {
	slug := c.Param("slug")
	var news dto.EditNewsRequest
	if err := c.ShouldBind(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Validator.Validate(&news)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.s.UpdateNews(slug, news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berita berhasil diperbarui",
	})
}

func (h *NewsControllers) DeleteNews(c *gin.Context) {
	slug := c.Param("slug")

	err := h.s.DeleteNews(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berita berhasil dihapus",
	})
}
