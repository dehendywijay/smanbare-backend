package config

import "os"

type Config struct {
	JWTAccessSecret string
	DATABASE_URL    string
	ADDRESS_REDIS   string
	USERNAME_REDIS  string
	PASSWORD_REDIS  string
}

func LoadConfig() *Config {
	return &Config{
		JWTAccessSecret: os.Getenv("JWT_ACCESS_SECRET"),
		DATABASE_URL:    os.Getenv("DATABASE_URL"),
		ADDRESS_REDIS:   os.Getenv("ADDRESS_REDIS"),
		USERNAME_REDIS:  os.Getenv("USERNAME_REDIS"),
		PASSWORD_REDIS:  os.Getenv("PASSWORD_REDIS"),
	}
}