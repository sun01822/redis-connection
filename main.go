package main

import (
	"github.com/sun01822/redis-connection/config"
	"github.com/sun01822/redis-connection/conn"
	"fmt"
)

func main() {
	config.LoadConfig()
	conn.ConnectRedis()

	client := conn.NewRedisClient()

	err := client.Set("appKey", "ea08dcb5-82ef-4738-bff4-3e3604b4d1f5", 0)
	if err!= nil {
        panic(err)
    }

	value, err := client.Get("appKey")
	if err!= nil {
        panic(err)
    }
	fmt.Println("Get AppKey: ", value)
}