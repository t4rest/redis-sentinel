package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")

	redis, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = redis.client.Set("test", 13, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := redis.client.Exists("test").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("res", res)

	t, err := redis.client.Get("test").Int64()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("t", t)

	err = redis.client.Del("test2").Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	res2, err := redis.client.Exists("test2").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("res2", res2)

	redis.Close()
}
