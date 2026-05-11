package controllers

import (
	"gin-app/models"
	"gin-app/services"
	"gin-app/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateKepalaSekolah(c *gin.Context) {
	name := c.PostForm("name")
	content := c.PostForm("content")

	if name == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and content are required"})
		return
	}
	
	fileBytes, objectPath, contentType, err := utility.ProcessImageUploadKepala(c, "image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publicURL, err := services.UploadToSupabase("kepala_sekolah", objectPath, contentType, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	kepalaSekolah := models.KepalaSekolah{
		Name:   name,
		Content: content,
		Foto: publicURL,
	}

	result, err := services.CreateKepalaSekolah(kepalaSekolah)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func EditKepalaSekolah(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	content := c.PostForm("content")

	if name == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and content are required"})
		return
	}
	
	fileBytes, objectPath, contentType, err := utility.ProcessImageUploadKepala(c, "image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publicURL, err := services.UploadToSupabase("kepala_sekolah", objectPath, contentType, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	kepalaSekolah := models.KepalaSekolah{
		Name:   name,
		Content: content,
		Foto: publicURL,
	}

	result, err := services.EditKepalaSekolah(id, kepalaSekolah)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}