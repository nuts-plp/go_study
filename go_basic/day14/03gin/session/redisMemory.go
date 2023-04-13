package session

import "github.com/go-redis/redis"

type RedisSession struct {
	sessionID  string
	sessionMap map[string]interface{}
	pool       *redis.Pool
}

func NewRedisMemory(id string,pool *redis.Pool)
