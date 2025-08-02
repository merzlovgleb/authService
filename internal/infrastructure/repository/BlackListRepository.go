package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ BlackListRepository = (*repository)(nil)

type BlackListRepository interface {
	AddToBlacklist(ctx context.Context, token string) error

	IsTokenBlacklisted(ctx context.Context, token string) (bool, error)
}
type repository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRepository(client *redis.Client, ttl time.Duration) BlackListRepository {
	return &repository{
		client: client,
		ttl:    ttl,
	}
}

func (repo *repository) AddToBlacklist(ctx context.Context, token string) error {

	return repo.client.SetEX(ctx, "blacklist:"+token, "1", repo.ttl).Err()
}

func (repo *repository) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {

	exists, err := repo.client.Exists(ctx, "blacklist:"+token).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}
