package worker

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func WorkerRedisFib(index string) (int64, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	val, err := redisClient.Get(ctx, index).Result()

	if err == redis.Nil {
		index64, _ := strconv.ParseInt(index, 10, 64)
		val64 := fib(index64)
		if err := redisClient.Set(ctx, index, val64, 0).Err(); err != nil {
			return val64, errors.New("redis saving error")
		}
	}
	if err != nil {
		return 0, errors.New("redis error")
	}
	val64, _ := strconv.ParseInt(val, 10, 64)
	return val64, nil
}
func fib(index int64) int64 {
	if index < 2 {
		return 1
	}
	return fib(index-1) + fib(index-2)
}

// redisClient
// sub.on('message', (channel, message) => {
//   redisClient.hset('values', message, fib(parseInt(message)));
// });
// sub.subscribe('insert');
