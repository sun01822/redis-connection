package conn

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/sun01822/redis-connection/config"
)


var client *redis.Client

func ConnectRedis(){
	conf := config.Redis()

	fmt.Println("connecting to redis at ", conf.Host, ":" + conf.Port, "...")

	client = redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
        Password: conf.Password,
        DB:       conf.DB,
	})

	if _, err := client.Ping().Result(); err != nil {
		fmt.Println("redis connection error: ", err)
		panic(err)
    } 
	
	fmt.Println("redis connected successfully")
}

type RedisClient struct {}

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (c *RedisClient) Get(key string) (string, error) {
    return client.Get(key).Result()
}

func (c *RedisClient) Set(key string, value string, expiration time.Duration) error {
    return client.Set(key, value, expiration).Err()
}

func (c *RedisClient) Del(keys ...string) (int, error) {
    err := client.Del(keys...)
	if err!= nil {
        return 0, fmt.Errorf("error deleting")
    }
	return len(keys), nil
}