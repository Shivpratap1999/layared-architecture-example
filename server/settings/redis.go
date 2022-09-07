package settings

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

// all credentials and configurations should be in env or yaml file 
func ConfigureRedis() (*redis.Client, error) {
	redisHost :=  os.Getenv("REDIS_HOST") //reading env
	redisPort := os.Getenv("REDIS_PORT")
	redisPassWord := os.Getenv("REDIS_PASSWORD")

	

	redisDB := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassWord,
		DB:       0,
	})

	// verifying redis connection
	// _, err := rdb.Ping(context.Background()).Result()
	result, err := redisDB.Ping().Result()
	if err != nil {
		return nil, err
	}

	log.Println("Redis Server Connected and Saying :", result)
	return redisDB, nil
}
