package conn

import (
	"github.com/go-redis/redis"
	"github.com/sun01822/redis-connection/config"
	"fmt"
)


var client *redis.Client

func ConnectRedis(){
	conf := config.Redis()

	fmt.Print("connecting to redis at ", conf.Host, ":" + conf.Port, "...")

	client = redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
        Password: conf.Password,
        DB:       conf.DB,
	})

	if _, err := client.Ping().Result(); err != nil {
		fmt.Print("redis connection error: ", err)
		panic(err)
    } 
	
	fmt.Print("redis connected successfully")
}