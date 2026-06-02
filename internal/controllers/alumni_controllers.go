package controllers

import (
	"gin-app/internal/dto"
	"gin-app/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AlumniControllers struct {
	s *services.AlumniServices
}

func NewAlumniControllers(s *services.AlumniServices) *AlumniControllers {
	return &AlumniControllers{
		s: s,
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

	err := h.s.CreateAlumni(c, &alumni)
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

	file := alumni.Foto
	fileName := file.Filename
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Foto tidak boleh kosong"})
		return
	}

	err := h.s.UpdateAlumni(c, id, alumni, fileName)
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
