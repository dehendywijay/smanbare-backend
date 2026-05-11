package controllers

import (
	"gin-app/models"
	"gin-app/services"
	"gin-app/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNews(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	kategori := c.PostForm("kategori")

	if title == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title and content are required"})
		return
	}

	fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publicURL, err := services.UploadToSupabase("news_thumbnail", objectPath, contentType, fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	slug := utility.MakeSlug(title)

	news := models.News{
		Title:     title,
		Kategori:  kategori,
		Content:   content,
		Thumbnail: publicURL,
		Slug:      slug,
	}

	result, err := services.CreateNews(news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berita berhasil dibuat",
		"data":    result,
	})
}

func GetNews(c *gin.Context) {
	result, err := services.GetNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetNewsByID(c *gin.Context) {
	slug := c.Param("slug")

	result, err := services.GetNewsByID(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateNews(c *gin.Context) {
	slug := c.Param("slug")
	title := c.PostForm("title")
	content := c.PostForm("content")
	kategori := c.PostForm("kategori")

	if title == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title and content are required"})
		return
	}

	news := models.News{
		Title:   title,
		Content: content,
		Slug:    utility.MakeSlug(title),
		Kategori: kategori,
	}

	file, _ := c.FormFile("image")
	if file != nil {
		oldObjectPath, err := services.GetFotoNews(slug)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get old thumbnail: " + err.Error()})
			return
		}

		fileBytes, objectPath, contentType, err := utility.ProcessImageUpload(c, "image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		oldPath := utility.ExtractObjectPath(oldObjectPath, "news_thumbnail")
		publicURL, err := services.UpdateSupabaseFile("news_thumbnail", oldPath, objectPath, contentType, fileBytes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		news.Thumbnail = publicURL
	}

	result, err := services.UpdateNews(slug, news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berita berhasil diperbarui",
		"data":    result,
	})
}

func DeleteNews(c *gin.Context) {
	slug := c.Param("slug")

	foto, err := services.DeleteNews(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fotopath := utility.ExtractObjectPath(foto, "news_thumbnail")

	err = services.DeleteFromSupabase("news_thumbnail", fotopath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Menhapus Foto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berita berhasil dihapus",
	})
}

