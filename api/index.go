package handler

import (
	"gin-app/config"
	"gin-app/middleware"
	"gin-app/routes"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	once   sync.Once
)

func setupRouter() {
	gin.SetMode(gin.ReleaseMode)

	config.ConnectDB()

	// err := config.DB.AutoMigrate(
	// 	&models.News{},
	// 	&models.Admin{},
	// 	&models.Guru{},
	// 	&models.KepalaSekolah{},
	// 	&models.Eskul{},
	// 	&models.Alumni{},
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	routes.NewsRoute(r)
	routes.AuthRoute(r)
	routes.GuruRoute(r)
	routes.EskulRoute(r)
	routes.AlumniRoute(r)

	router = r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(setupRouter)
	router.ServeHTTP(w, r)
}