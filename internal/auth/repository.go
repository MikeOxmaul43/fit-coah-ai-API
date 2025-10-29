package auth

import (
	"context"
	"sportTrackerAPI/redisDb"
	"time"
)

type Repository struct {
	RedisDataBase *redisDb.RDb
}

func NewAuthRepository(rdb *redisDb.RDb) *Repository {
	return &Repository{RedisDataBase: rdb}
}
func (repo *Repository) Set(email, refreshToken string, expireAt time.Duration) error {
	var ctx = context.Background()
	err := repo.RedisDataBase.Client.Set(ctx, email, refreshToken, expireAt).Err()
	if err != nil {
		return err
	}
	return nil
}
