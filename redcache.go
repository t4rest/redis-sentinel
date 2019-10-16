package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Redis redis struct
type Redis struct {
	client *redis.Client
}

// New create new pool redis connections
func New() (*Redis, error) {

	fo := &redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{":26379"},
		OnConnect: func(conn *redis.Conn) error {
			return conn.Ping().Err()
		},
		Password:     "123",
		DB:           1,
		DialTimeout:  time.Second * 15,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		PoolSize:     5,
		PoolTimeout:  time.Second * 15,
	}

	rc := redis.NewFailoverClient(fo)

	rds := &Redis{
		client: rc,
	}
	err := rc.Ping().Err()

	return rds, err
}

// Close pool of connections
func (rds *Redis) Close() {
	err := rds.client.Close()
	if err != nil {
		fmt.Println(err)
	}
}
