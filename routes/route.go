package routes

import (
	"gin-app/controllers"
	"gin-app/middleware"
	"github.com/gin-gonic/gin"
)

func NewsRoute(r *gin.Engine) {
	news := r.Group("/api/news")
	{
		news.GET("", controllers.GetNews)
		news.GET("/", controllers.GetNews)
		news.GET("/:slug", controllers.GetNewsByID)
		
	}
	news.Use(middleware.AuthMiddleware())
	{
		news.POST("", controllers.CreateNews)
		news.PUT("/:slug", controllers.UpdateNews)
		news.DELETE("/:slug", controllers.DeleteNews)
		
	}

	detail := r.Group("/api/detail")
	{
		detail.POST("/kepala-sekolah", controllers.CreateKepalaSekolah)
		detail.PUT("/kepala-sekolah/:id", controllers.EditKepalaSekolah)
	}
}

func AuthRoute(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.LoginAdmin)
		auth.POST("/refresh", controllers.RefreshToken)
	}
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/logout", controllers.LogoutAdmin)
	}
}

func GuruRoute(r *gin.Engine) {
	guru := r.Group("/api/guru")
	{
		guru.GET("", controllers.GetGuru)
		guru.GET("/kepala/:id", controllers.GetKepalaByID)
		
	}
	guru.Use(middleware.AuthMiddleware())
	{
		guru.POST("", controllers.CreateGuru)
		guru.PUT("/:id", controllers.EditGuru)
		guru.POST("/kepala", controllers.CreateKepala)
		guru.PUT("/kepala/:id", controllers.EditKepala)
		guru.DELETE("/:id", controllers.DeleteGuru)
	}
}

func EskulRoute(r *gin.Engine) {
	eskul := r.Group("/api/eskul")
	{
		eskul.GET("", controllers.GetEskul)
		eskul.GET("/:slug", controllers.GetEskulByID)
	}
	eskul.Use(middleware.AuthMiddleware())
	{
		eskul.POST("", controllers.CreateEskul)
		eskul.PUT("/:slug", controllers.EditEskul)
		eskul.DELETE("/:slug", controllers.DeleteEskul)
	}
}

func AlumniRoute(r *gin.Engine) {
	alumni := r.Group("/api/alumni")
	{
		
		alumni.GET("", controllers.GetAllAlumni)
		
	}
	alumni.Use(middleware.AuthMiddleware())
	{
		alumni.PUT("/:id", controllers.UpdateAlumni)
		alumni.DELETE("/:id", controllers.DeleteAlumni)
		alumni.POST("", controllers.CreateAlumni)
	}
}