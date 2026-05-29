package config

import (
	"gin-app/internal/controllers"
	"gin-app/internal/repository"
	"gin-app/internal/services"

	"gorm.io/gorm"
)

type App struct {
	AuthController   *controllers.AuthControllers
	AlumniController *controllers.AlumniControllers
	EskulController  *controllers.EskulControllers
	GuruController   *controllers.GuruControllers
	NewsController   *controllers.NewsControllers
}

func BootstrapApp(db *gorm.DB) *App {

	authRepo := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthControllers(authService)

	alumniRepo := repository.NewAlumniRepository(db)
	alumniService := services.NewAlumniService(alumniRepo)
	alumniController := controllers.NewAlumniControllers(alumniService)

	eskulRepo := repository.NewEskulRepository(db)
	eskulService := services.NewEskulService(eskulRepo)
	eskulController := controllers.NewEskulControllers(eskulService)

	guruRepo := repository.NewGuruRepository(db)
	guruService := services.NewGuruService(guruRepo)
	guruController := controllers.NewGuruControllers(guruService)

	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)
	newsController := controllers.NewNewsControllers(newsService)

	return &App{
		AuthController:   authController,
		AlumniController: alumniController,
		EskulController:  eskulController,
		GuruController:   guruController,
		NewsController:   newsController,
	}
}