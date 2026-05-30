package api	

import (
	"gin-app/internal/config"
	"gin-app/internal/middleware"
	"gin-app/internal/routes"
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

	app := config.BootstrapApp(config.DB)

	routes.SetupRoutes(r, app)

	router = r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(setupRouter)
	router.ServeHTTP(w, r)
}