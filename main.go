package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisURL := fmt.Sprintf("%s:%s", redisHost, redisPort)
	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	err := client.Set("mykey", "Hello Redis", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := client.Get("mykey").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value from Redis:", val)
}
