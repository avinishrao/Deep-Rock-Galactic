package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use the appropriate address
		Password: "",               // no password set by default
		DB:       0,                // use default DB
	})

	// Check the connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong) // Output: PONG

	// Set the keys
	fmt.Println("----------------------Create----------------------")
	err = rdb.Set(ctx, "key1", "oil", 0).Err()
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, "key2", "dystrum", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Added new keys & values")

	// Get the key
	fmt.Println("----------------------Read------------------------")
	val1, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1:", val1)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key2:", val2)

	// Optionally, close the connection when your application exits
	err = rdb.Close()
	if err != nil {
		panic(err)
	}
}