package utils

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var Store *session.Store

func InitRedisSession() {
	redisStore := redis.New()

	Store = session.New(session.Config{
		Storage: redisStore,
	})
}
