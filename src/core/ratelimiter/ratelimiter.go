package ratelimiter

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	client     *redis.Client
	limit      int
	window     time.Duration
}


func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
        panic(err)
    }
	client := redis.NewClient(opts)

	return &RateLimiter{
		client: client,
		limit:  limit,
		window: window,
	}
}


func (rl *RateLimiter) Allow(ctx context.Context, key string) (bool, error) {
	cacheKey := key
	pipe := rl.client.TxPipeline()
	counter := pipe.Incr(ctx, cacheKey)
	pipe.Expire(ctx, cacheKey, rl.window)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}
	return counter.Val() <= int64(rl.limit), nil
}