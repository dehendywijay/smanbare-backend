package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisConnect(cfg *Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.ADDRESS_REDIS,
		Username: cfg.USERNAME_REDIS,
		Password: cfg.PASSWORD_REDIS,
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Gagal konek Redis: %v", err)
		panic(err)
	}

	log.Println("Koneksi ke Redis berhasil")
	return rdb, nil

}
