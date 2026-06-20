package handler

import (
	"gin-app/config"
	"gin-app/middleware"
	"gin-app/routes"
	"log"
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

	cfg := config.LoadConfig()

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	redis, err := config.RedisConnect(cfg)
	if err != nil {
		log.Fatal("Gagal konek ke Redis:", err)
	}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	app := config.BootstrapApp(db, cfg, redis)

	routes.SetupRoutes(r, app)

	router = r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(setupRouter)
	router.ServeHTTP(w, r)
}
