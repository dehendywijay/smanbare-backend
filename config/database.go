package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if os.Getenv("VERCEL") == "" {
		_ = godotenv.Load()
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		fmt.Println("DATABASE_URL belum diset")
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		fmt.Println("Gagal konek database:", err)
		os.Exit(1)
	}

	DB = db
	fmt.Println("Database connected ✅")
}