package server

import (
	"gin-app/config"
	"gin-app/middleware"
	"gin-app/models"
	"gin-app/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
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
		log.Fatal("Gagal migrasi:", err)
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

	return r
}