package main

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

const redisHost = "redis://localhost:6380"
const defaultChannel = "sample"
const defaultMessage = "hello"

func newRedisPool(host string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(host)
			if err != nil {
				log.Printf("ERROR: failed initializing the redis pool: %s", err.Error())
				os.Exit(1)
			}
			return c, err
		},
	}
}

func main() {
	// Init a Redis pool and fetch a connection
	rPool := newRedisPool(redisHost)
	redisConn := rPool.Get()
	defer redisConn.Close()

	// Init channel
	channel := defaultChannel
	if len(os.Args) > 1 {
		channel = os.Args[1]
	}

	// Init message
	message := defaultMessage
	if len(os.Args) > 2 {
		message = os.Args[2]
	}

	redisConn.Do("PUBLISH", channel, message)
}
