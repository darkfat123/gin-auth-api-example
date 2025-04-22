package redis

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

func SetData(key string, value interface{}, expiration time.Duration) error {
	err := Rdb.Set(context.TODO(), key, value, expiration).Err()
	return err
}

func GetData(key string) (string, error) {
	val, err := Rdb.Get(context.TODO(), key).Result()
	return val, err
}

func DeleteData(key string) error {
	err := Rdb.Del(context.TODO(), key).Err()
	return err
}
