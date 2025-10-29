package redisDb

import (
	"github.com/redis/go-redis/v9"
	"sportTrackerAPI/internal/config"
)

type RDb struct {
	*redis.Client
}

func NewRDb(cfg *config.Config) *RDb {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Rdb.Address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RDb{rdb}
}
