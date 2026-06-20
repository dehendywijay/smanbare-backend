package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB(cfg *Config) (*gorm.DB, error) {
	start := time.Now()
    dsn := cfg.DATABASE_URL

    if dsn == "" {
        log.Println("DATABASE belum diset")
    }

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal("Gagal konek database:", err, "ss", dsn)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Gagal mengambil database instance:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("Database connected", time.Since(start))

	return db, nil
}