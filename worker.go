package worker

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func WorkerRedisFib(index int64, redisUrl string) (int64, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisUrl, //"localhost:6379"
		Password: "",       // no password set
		DB:       0,        // use default DB
	})
	pong, err := redisClient.Ping(ctx).Result()
	fmt.Println(pong, err)
	index_string := strconv.FormatInt(index, 10)
	val, err := redisClient.Get(ctx, index_string).Result()

	if err == redis.Nil {
		val := fib(index)
		if err := redisClient.Set(ctx, index_string, val, 0).Err(); err != nil {
			return val, errors.New("redis saving error")
		}
		return val, nil
	}
	if err != nil {
		return 0, errors.New(err.Error())
	}
	val64, _ := strconv.ParseInt(val, 10, 64)
	return val64, nil
}
func fib(index int64) int64 {
	if index <= 2 {
		return 1
	}
	return fib(index-1) + fib(index-2)
}
