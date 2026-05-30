package routes

import (
	"gin-app/config"
	"gin-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, app *config.App) {

	// AUTH
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", app.AuthController.LoginAdmin)
		auth.POST("/refresh", app.AuthController.RefreshToken)
	}

	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/logout", app.AuthController.LogoutAdmin)
	}

	// NEWS
	news := r.Group("/api/news")
	{
		news.GET("", app.NewsController.GetNews)
		news.GET("/", app.NewsController.GetNews)
		news.GET("/:slug", app.NewsController.GetNewsByID)
	}

	news.Use(middleware.AuthMiddleware())
	{
		news.POST("", app.NewsController.CreateNews)
		news.PUT("/:slug", app.NewsController.UpdateNews)
		news.DELETE("/:slug", app.NewsController.DeleteNews)
	}

	
	// GURU
	guru := r.Group("/api/guru")
	{
		guru.GET("", app.GuruController.GetGuru)
		guru.GET("/kepala/:id", app.GuruController.GetKepalaByID)
	}

	guru.Use(middleware.AuthMiddleware())
	{
		guru.POST("", app.GuruController.CreateGuru)
		guru.PUT("/:id", app.GuruController.EditGuru)
		// guru.POST("/kepala", app.GuruController.CreateKepala)
		guru.PUT("/kepala/:id", app.GuruController.EditKepala)
		guru.DELETE("/:id", app.GuruController.DeleteGuru)
	}

	// ESKUL
	eskul := r.Group("/api/eskul")
	{
		eskul.GET("", app.EskulController.GetEskul)
		eskul.GET("/:slug", app.EskulController.GetEskulByID)
	}

	eskul.Use(middleware.AuthMiddleware())
	{
		eskul.POST("", app.EskulController.CreateEskul)
		eskul.PUT("/:slug", app.EskulController.EditEskul)
		eskul.DELETE("/:slug", app.EskulController.DeleteEskul)
	}

	// ALUMNI
	alumni := r.Group("/api/alumni")
	{
		alumni.GET("", app.AlumniController.GetAllAlumni)
	}

	alumni.Use(middleware.AuthMiddleware())
	{
		alumni.PUT("/:id", app.AlumniController.UpdateAlumni)
		alumni.DELETE("/:id", app.AlumniController.DeleteAlumni)
		alumni.POST("", app.AlumniController.CreateAlumni)
	}
}