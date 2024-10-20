package main

import (
	"github.com/sun01822/redis-connection/config"
	"github.com/sun01822/redis-connection/conn"
)

func main() {
	config.LoadConfig()
	conn.ConnectRedis()
}