package controllers

import (
	"gin-app/models"
	"gin-app/services"
	"net/http"
	"gin-app/utility"

	"github.com/gin-gonic/gin"
)

func CreateEskul(c *gin.Context) {
	nama := c.PostForm("nama")
	pembina := c.PostForm("pembina")
	jadwal := c.PostForm("jadwal")
	prestasi := c.PostForm("prestasi")
	tujuan := c.PostForm("tujuan")

	slug := utility.MakeSlug(nama)

	

	fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal memproses gambar: " })
		return
	}
	
	publicURL, err := services.UploadToSupabase("eskul", objectPath, contentType, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengunggah gambar " })
		return
	}

	eskul := models.Eskul{
		Nama:    nama,
		Pembina: pembina,
		Jadwal:  jadwal,
		Prestasi: prestasi,
		Foto:    publicURL,
		Tujuan: tujuan,
		Slug:    slug,
	}


	_, err = services.CreateEskul(eskul)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Data berhasil dibuat")
	}


func GetEskul(c *gin.Context) {
	eskul, err := services.GetEskul()	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, eskul)
}

func GetEskulByID(c *gin.Context) {
	id := c.Param("slug")
	eskul, err := services.GetEskulByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eskul)
}

func EditEskul(c *gin.Context) {
	id := c.Param("slug")
	nama := c.PostForm("nama")
	pembina := c.PostForm("pembina")
	jadwal := c.PostForm("jadwal")
	prestasi := c.PostForm("prestasi")
	tujuan := c.PostForm("tujuan")

	newSlug := utility.MakeSlug(nama)

	

	eskul := models.Eskul{
		Nama:    nama,
		Pembina: pembina,
		Jadwal:  jadwal,
		Prestasi: prestasi,
		Tujuan: tujuan,
		Slug: newSlug,
	}

	file, _ := c.FormFile("foto")
	if file != nil {
		oldObjectPath, err := services.GetFotoEskul(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data eskul " })
			return
		}

		oldFoto := utility.ExtractObjectPath(oldObjectPath, "eskul")

		fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "foto")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		publicURL, err := services.UpdateSupabaseFile("eskul", oldFoto , objectPath, contentType, fileBytes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		eskul.Foto = publicURL
	}


	_ , err := services.EditEskul(id, eskul)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Data berhasil diperbarui")
}

func DeleteEskul(c *gin.Context) {
	id := c.Param("slug")

	foto, err := services.GetFotoEskul(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data eskul " })
		return
	}
	
	if foto != "" {
		fotopath := utility.ExtractObjectPath(foto, "eskul")
		err = services.DeleteFromSupabase("eskul",fotopath)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "Gagal Menhapus Foto"})
			return
		}
	}

	err = services.DeleteEskul(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Data berhasil dihapus")
}