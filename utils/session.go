package utils

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var Store *session.Store

func InitRedisSession() {
	redisStore := redis.New(redis.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Username: "",
		Password: "",
		Database: 0,
		Reset:    false,
	})
	Store = session.New(session.Config{
		Storage: redisStore,
	})

}
