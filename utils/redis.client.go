package utils

import (
	"context"
	"crypto/tls"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(ctx context.Context, addr string, username string, password string, db int, tls *tls.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:      addr,
		Username:  username,
		Password:  password,
		DB:        db,
		TLSConfig: tls,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
