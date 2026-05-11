package main

import (
	"gin-app/config"
	"gin-app/models"
	"gin-app/routes"
	"os"
	"gin-app/middleware"
	"github.com/gin-gonic/gin"

)


func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())


	config.ConnectDB()
	config.DB.AutoMigrate(&models.News{}, &models.Admin{}, &models.Guru{}, &models.KepalaSekolah{}, &models.Eskul{}, &models.Alumni{})
	
	routes.NewsRoute(r)
	routes.AuthRoute(r)
	routes.GuruRoute(r)
	routes.EskulRoute(r)
	routes.AlumniRoute(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r.Run(":" + port)
}