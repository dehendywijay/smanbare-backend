package controllers

import (
	"gin-app/models"
	"gin-app/utility"
	"gin-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGuru(c *gin.Context) {
	nama := c.PostForm("nama")
	jabatan := c.PostForm("jabatan")

	if nama == "" || jabatan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama and jabatan are required"})
		return
	}
	fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publicURL, err := services.UploadToSupabase("guru", objectPath, contentType, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengunggah gambar " })
		return
	}

	guru := models.Guru{
		Nama:    nama,
		Jabatan: jabatan,
		Foto:    publicURL,
	}

	_ , err = services.CreateGuru(guru)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data guru " })
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dibuat",
	})
}

func GetGuru(c *gin.Context) {
	guru, err := services.GetGuru()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guru)
}

func EditGuru(c *gin.Context) {
	id := c.Param("id")

	nama := c.PostForm("nama")
	jabatan := c.PostForm("jabatan")

	if nama == "" || jabatan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama and jabatan are required"})
		return
	}
	
	guru := models.Guru{
		Nama:    nama,
		Jabatan: jabatan,
	}

	file, _ := c.FormFile("foto")
	if file != nil {
		oldObjectPath, err := services.GetFotoGuru(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data guru " })
			return
		}

		oldFoto := utility.ExtractObjectPath(oldObjectPath, "guru")

		fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "foto")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		publicURL, err := services.UpdateSupabaseFile("guru", oldFoto , objectPath, contentType, fileBytes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		guru.Foto = publicURL
	}


	_ , err := services.EditGuru(id, guru)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data guru " })
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil diupdate",
	})
}

func DeleteGuru(c *gin.Context) {
	id := c.Param("id")	

	foto, err := services.GetFotoGuru(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data guru " })
		return
	}
	
	if foto != "" {
		fotopath := utility.ExtractObjectPath(foto, "guru")
		err = services.DeleteFromSupabase("guru",fotopath)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "Gagal Menhapus Foto"})
			return
		}
	}

	err = services.DeleteGuru(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dihapus",
	})
}
	

func EditKepala(c *gin.Context) {
	id := c.Param("id")

	nama := c.PostForm("nama")
	content := c.PostForm("content")
	
	kepalaSekolah := models.KepalaSekolah{
		Name:    nama,
		Content: content,
	}

	file, _ := c.FormFile("foto")
	if file != nil {
		oldObjectPath, err := services.GetFotoKepala(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data kepala sekolah " })
			return
		}

		oldFoto := utility.ExtractObjectPath(oldObjectPath, "kepala")

		fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "foto")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		publicURL, err := services.UpdateSupabaseFile("kepala", oldFoto , objectPath, contentType, fileBytes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		kepalaSekolah.Foto = publicURL
	}


	_ , err := services.EditKepalaSekolah(id, kepalaSekolah)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data kepala sekolah " })
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil diupdate",
	})
}

func CreateKepala(c *gin.Context) {
	nama := c.PostForm("nama")
	content := c.PostForm("content")

	fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publicURL, err := services.UploadToSupabase("kepala", objectPath, contentType, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengunggah gambar " })
		return
	}

	kepalaSekolah := models.KepalaSekolah{
		Name:    nama,
		Content: content,
		Foto:    publicURL,
	}

	_ , err = services.CreateKepalaSekolah(kepalaSekolah)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data kepala sekolah " })
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dibuat",
	})
}

func GetKepalaByID(c *gin.Context) {
	id := c.Param("id")

	kepalaSekolah, err := services.GetKepalaByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data kepala sekolah " })
		return
	}

	c.JSON(http.StatusOK, kepalaSekolah)
}
