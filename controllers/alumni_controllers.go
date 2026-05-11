package controllers

import (
	"gin-app/models"
	"gin-app/services"
	"gin-app/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAlumni(c *gin.Context) {
	alumni, err := services.GetAllAlumni()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve alumni"})
		return
	}
	c.JSON(http.StatusOK, alumni)
}

func CreateAlumni(c *gin.Context) {
	nama := c.PostForm("nama")
	universitas := c.PostForm("universitas")
	tahun := c.PostForm("tahun")
	
	fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process image: " + err.Error()})
		return
	}

	publicURL, err := services.UploadToSupabase("alumni", objectPath, contentType, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image: " + err.Error()})
		return
	}

	alumni := models.Alumni{
		Nama:        nama,
		Foto:        publicURL,
		Universitas: universitas,
		Tahun: tahun,
	}

	_, err = services.CreateAlumni(alumni)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create alumni: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Data berhasil dibuat")
}

func UpdateAlumni(c *gin.Context) {	
	id := c.Param("id")
	nama := c.PostForm("nama")
	universitas := c.PostForm("universitas")
	tahun := c.PostForm("tahun")

	alumni := models.Alumni{
			Nama:        nama,
			Universitas: universitas,
			Tahun: tahun,
		}

	file, _ := c.FormFile("foto")
	if file != nil {
		oldObjectPath, err := services.GetFotoAlumni(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get alumni data: " + err.Error()})
			return
		}

		oldFoto := utility.ExtractObjectPath(oldObjectPath, "alumni")

		fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "foto")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		publicURL, err := services.UpdateSupabaseFile("alumni", oldFoto , objectPath, contentType, fileBytes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		alumni.Foto = publicURL
	} 
		

		_, err := services.UpdateAlumni(id, alumni)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data guru " })
			return
		}

		c.JSON(http.StatusOK, "Data berhasil diupdate")
	}


func DeleteAlumni(c *gin.Context) {
	id := c.Param("id")

	foto, err := services.GetFotoAlumni(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data alumni " })
		return
	}
	
	if foto != "" {
		fotopath := utility.ExtractObjectPath(foto, "alumni")
		err = services.DeleteFromSupabase("alumni",fotopath)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "Gagal Menhapus Foto"})
			return
		}
	}

	err = services.DeleteAlumni(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete alumni: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Data berhasil dihapus")
}