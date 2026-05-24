package main

import (
	"gin-app/config"
	"gin-app/middleware"
	"gin-app/models"
	"gin-app/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	
	config.ConnectDB()

	err := config.DB.AutoMigrate(
		&models.News{},
		&models.Admin{},
		&models.Guru{},
		&models.KepalaSekolah{},
		&models.Eskul{},
		&models.Alumni{},
	)
	if err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	routes.NewsRoute(r)
	routes.AuthRoute(r)
	routes.GuruRoute(r)
	routes.EskulRoute(r)
	routes.AlumniRoute(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("Server running on port", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}