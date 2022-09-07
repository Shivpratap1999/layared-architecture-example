package jsonRepo

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

type redisStorer struct {
	redisClint *redis.Client
}

func NewredisStorer(R *redis.Client) *redisStorer {
	return &redisStorer{redisClint: R}
}
func (r redisStorer) StoreNewObject(key string, value interface{}, expiry time.Duration) error {
	err := r.redisClint.Set(key, value, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r redisStorer) DeleteObject(key string) error {
	_, err := r.redisClint.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}
func (r redisStorer) IsObjectExist(key string) bool {
	_, err := r.redisClint.Get(key).Result()
	if err != nil {
		log.Printf("[jsonRepo]error when cheching object in Redis  %s\n", err)
		return false
	}
	return true
}
