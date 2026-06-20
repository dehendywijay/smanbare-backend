package controllers

import (
	"gin-app/internal/dto"
	"gin-app/internal/services"
	"gin-app/internal/validator"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AlumniControllers struct {
	s *services.AlumniServices
	Validator validator.CustomValidator
}

func NewAlumniControllers(s *services.AlumniServices, validator validator.CustomValidator) *AlumniControllers {
	return &AlumniControllers{
		s: s,
		Validator: validator,
	}
}

func (h *AlumniControllers) GetAllAlumni(c *gin.Context) {
	hasil, err := h.s.GetAllAlumni()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Mengambil data Alumni: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, hasil)
}

func (h *AlumniControllers) CreateAlumni(c *gin.Context) {
	var alumni dto.AlumniCreateRequest
	if err := c.ShouldBind(&alumni); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data yang dimasukkan tidak sesuai: " + err.Error()})
		return
	}

	err := h.Validator.Validate(&alumni)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.s.CreateAlumni(&alumni)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data alumni: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dibuat"})
}

func (h *AlumniControllers) UpdateAlumni(c *gin.Context) {
	id := c.Param("id")
	var alumni dto.AlumniEditRequest
	if err := c.ShouldBind(&alumni); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data yang dimasukkan tidak sesuai: " + err.Error()})
		return
	}

	err := h.Validator.Validate(&alumni)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.s.UpdateAlumni(id, alumni)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data alumni: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func (h *AlumniControllers) DeleteAlumni(c *gin.Context) {
	id := c.Param("id")

	err := h.s.DeleteAlumni(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data alumni: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
