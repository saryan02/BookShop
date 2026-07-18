package db

import (
	"BookShop/internal/config"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Redis(cfg *config.Config) (*redis.Client, error) {

	opt := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisHost,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	err := opt.Ping(context.Background()).Err()

	if err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return opt, nil
}
