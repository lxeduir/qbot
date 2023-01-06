package db

import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)

var redisDb *redis.Client

func InitClirnt() (err error) {
	addr := "123.249.92.218:6379"
	redisDb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "123456",
		DB:       8,
	})
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func Set(key string, value string, time time.Duration) (err error) {
	err = redisDb.Set(key, value, time).Err()
	if err != nil {
		return err
	}
	return nil
}
func Get(key string) (val string, err error) {
	val, err = redisDb.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
func Del(key string) (err error) {
	redisDb.Del(key)
	_, err = Get(key)
	if err == nil {
		e := errors.New("删除失败")
		return e
	}
	return nil
}
