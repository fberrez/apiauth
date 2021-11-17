package backend

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/juju/errors"
)

// NewRedis initializes a new redis client.
// When the client is created, it performs a ping on the redis server.
func NewRedis(addr, password string, ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Annotatef(err, "ping redis")
	}

	return client, nil
}
