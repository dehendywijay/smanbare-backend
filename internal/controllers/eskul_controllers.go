package controllers

import (
	"gin-app/internal/dto"
	"gin-app/internal/services"
	"gin-app/pkg/slug"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EskulControllers struct {
	s *services.EskulService
}

func NewEskulControllers(s *services.EskulService) *EskulControllers {
	return &EskulControllers{
		s: s,
	}
}

func (h *EskulControllers) CreateEskul(c *gin.Context) {
	var eskul dto.CreateEskulRequest

	if err := c.ShouldBind(&eskul); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eskul.Slug = slug.MakeSlug(eskul.Nama)

	err := h.s.CreateEskul(eskul, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Data berhasil dibuat")
}

func (h *EskulControllers) GetEskul(c *gin.Context) {
	eskul, err := h.s.GetEskul()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, eskul)
}

func (h *EskulControllers) GetEskulByID(c *gin.Context) {
	id := c.Param("slug")
	eskul, err := h.s.GetEskulByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eskul)
}

func (h *EskulControllers) EditEskul(c *gin.Context) {
	id := c.Param("slug")
	var eskul dto.EditEskulRequest

	if err := c.ShouldBind(&eskul); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eskul.Slug = slug.MakeSlug(eskul.Nama)

	err := h.s.EditEskul(id, eskul, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Data berhasil diperbarui")
}

func (h *EskulControllers) DeleteEskul(c *gin.Context) {
	id := c.Param("slug")

	err := h.s.DeleteEskul(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Data berhasil dihapus")
}
