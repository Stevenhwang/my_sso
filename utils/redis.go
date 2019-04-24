package utils

import (
	"time"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "192.168.50.181:6379",
	Password: "",
	DB:       0,
	PoolSize: 20,
})

func RefreshToken(name, token string) string {
	err := client.Set(name, token, 10*time.Hour).Err()
	if err != nil {
		panic(err)
	}
	return "Refresh token success!"
}

func GetToken(name string) string {
	token, err := client.Get(name).Result()
	if err != nil {
		return name + " has no token!"
	}
	return token
}

func DelToken(name string) string {
	_, err := client.Get(name).Result()
	if err != nil {
		return name + " has no token!"
	}
	client.Del(name)
	return "Delete " + name + " token success!"
}
