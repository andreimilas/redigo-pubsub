package main

import (
	"log"
	"os"
	"sync"

	"github.com/gomodule/redigo/redis"
)

const redisHost = "redis://localhost:6380"
const defaultChannel = "sample"

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

	// Pub/Sub connection
	psConn := redis.PubSubConn{Conn: redisConn}
	if err := psConn.Subscribe(redis.Args{}.AddFlat(channel)...); err != nil {
		log.Printf("ERROR: failed initializing the redis pool: %s", err.Error())
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Start a goroutine to receive notifications from the server.
	go func() {
		for {
			switch v := psConn.Receive().(type) {
			case redis.Message:
				log.Printf("New Message: '%s' on channel '%s'", string(v.Data), v.Channel)
			case redis.Subscription:
				log.Printf("New Subscriber: '%v'", v)
			case error:
				log.Printf("Error: %+v", v)
				return
			}
		}
	}()
	wg.Wait()
}
