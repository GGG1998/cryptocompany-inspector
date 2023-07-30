package redis

import (
	"context"
	"cryptocompany-inspector/pkg/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	config *config.RedisConfigurator
	client *redis.Client
}

func (r *Redis) Client(ctx context.Context) (*redis.Client, error) {
	if r.client == nil {
		return nil, fmt.Errorf("redis client is nill")
	}
	pingCmd := r.client.Ping(ctx)
	_, err := pingCmd.Result()
	if err != nil {
		return nil, fmt.Errorf("redis client: %w", err)
	}

	return r.client, nil
}

func (r *Redis) LPush(ctx context.Context, key, val string) error {
	// Check client nil pointer
	_, err := r.Client(ctx)
	if err != nil {
		return err
	}

	lpushCommand := r.client.LPush(ctx, key, val)
	if _, err := lpushCommand.Result(); err != nil {
		return err
	}

	return nil
}

func NewRedis(config *config.RedisConfigurator) *Redis {
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.Port,
	})
	return &Redis{
		config: config,
		client: rdb,
	}
}
