package config

import (
	"gin-app/internal/controllers"
	"gin-app/internal/repository"
	"gin-app/internal/services"
	"gin-app/internal/validator"
	jwttoken "gin-app/pkg/jwt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	AuthController   *controllers.AuthControllers
	AlumniController *controllers.AlumniControllers
	EskulController  *controllers.EskulControllers
	GuruController   *controllers.GuruControllers
	NewsController   *controllers.NewsControllers
	Config           *Config
}

func BootstrapApp(db *gorm.DB, cfg *Config, redis *redis.Client) *App {
	validator := validator.NewCustomValidator()
	token := jwttoken.NewJWTService(cfg.JWTAccessSecret)

	authRepo := repository.NewAuthRepository(db)
	alumniRepo := repository.NewAlumniRepository(db)
	eskulRepo := repository.NewEskulRepository(db)
	guruRepo := repository.NewGuruRepository(db)
	newsRepo := repository.NewNewsRepository(db)
	tokenRepo := repository.NewTokenRepository(redis)


	authService := services.NewAuthService(authRepo, token, tokenRepo)
	alumniService := services.NewAlumniService(alumniRepo)
	eskulService := services.NewEskulService(eskulRepo)
	guruService := services.NewGuruService(guruRepo)
	newsService := services.NewNewsService(newsRepo)

	authController := controllers.NewAuthControllers(authService, validator, tokenRepo)
	alumniController := controllers.NewAlumniControllers(alumniService, validator)
	eskulController := controllers.NewEskulControllers(eskulService, validator)
	guruController := controllers.NewGuruControllers(guruService, validator)
	newsController := controllers.NewNewsControllers(newsService, validator)

	return &App{
		AuthController:   authController,
		AlumniController: alumniController,
		EskulController:  eskulController,
		GuruController:   guruController,
		NewsController:   newsController,
		Config:           cfg,
	}
}