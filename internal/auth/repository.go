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
	ctx := context.Background()
	err := repo.RedisDataBase.Client.Set(ctx, email, refreshToken, expireAt).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) Get(email string) (string, error) {
	ctx := context.Background()
	val, err := repo.RedisDataBase.Client.Get(ctx, email).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (repo *Repository) Delete(email string) error {
	ctx := context.Background()
	return repo.RedisDataBase.Client.Del(ctx, email).Err()
}
