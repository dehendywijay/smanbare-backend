package controllers

import (
	"gin-app/internal/dto"
	"gin-app/internal/services"
	"gin-app/internal/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GuruControllers struct {
	s         *services.GuruService
	Validator validator.CustomValidator
}

func NewGuruControllers(s *services.GuruService, validator validator.CustomValidator) *GuruControllers {
	return &GuruControllers{
		s:         s,
		Validator: validator,
	}
}

func (h *GuruControllers) CreateGuru(c *gin.Context) {
	var guru dto.GuruRequest
	if err := c.ShouldBind(&guru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Validator.Validate(&guru)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.s.CreateGuru(guru)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dibuat",
	})
}

func (h *GuruControllers) GetGuru(c *gin.Context) {
	guru, err := h.s.GetGuru()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guru)
}

func (h *GuruControllers) EditGuru(c *gin.Context) {
	id := c.Param("id")
	var guru dto.EditGuruRequest
	if err := c.ShouldBind(&guru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Validator.Validate(&guru)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.s.EditGuru(id, guru)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil diupdate",
	})
}

func (h *GuruControllers) DeleteGuru(c *gin.Context) {
	id := c.Param("id")

	err := h.s.DeleteGuru(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dihapus",
	})
}

func (h *GuruControllers) EditKepala(c *gin.Context) {
	id := c.Param("id")
	var kepalaSekolah dto.KepalaSekolahRequest
	if err := c.ShouldBind(&kepalaSekolah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if kepalaSekolah.Name == "" || kepalaSekolah.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and content are required"})
		return
	}

	err := h.s.EditKepala(id, kepalaSekolah)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data kepala sekolah "})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil diupdate",
	})
}

// func (h *GuruControllers) CreateKepala(c *gin.Context) {
// 	nama := c.PostForm("nama")
// 	content := c.PostForm("content")

// 	fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "foto")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	publicURL, err := h.s.UploadToSupabase("kepala", objectPath, contentType, fileBytes)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengunggah gambar " })
// 		return
// 	}

// 	kepalaSekolah := models.KepalaSekolah{
// 		Name:    nama,
// 		Content: content,
// 		Foto:    publicURL,
// 	}

// 	_ , err = h.s.CreateKepala(kepalaSekolah)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data kepala sekolah " })
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Data berhasil dibuat",
// 	})
// }

func (h *GuruControllers) GetKepalaByID(c *gin.Context) {
	id := c.Param("id")

	kepalaSekolah, err := h.s.GetKepalaByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data kepala sekolah "})
		return
	}

	c.JSON(http.StatusOK, kepalaSekolah)
}
