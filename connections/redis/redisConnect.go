package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

const (
	REDISCONNECTIONSTRING = "redis:6379"
	PASSWORD = ""
	DB = 0
)

var Ctx = context.Background()

func getClient(ctx context.Context) *redis.Client{
	log.Println("connecting ---to----redis----")

	client := redis.NewClient(&redis.Options{
		Addr:     REDISCONNECTIONSTRING,
		Password: PASSWORD,
		DB:       DB,
	})
	log.Println("connected")

	return client
}

func runRedis() {

	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping(Ctx).Result()
	fmt.Println(pong, err)

	//err = client.Set("name", "Elliot", 0).Err()
	//val, err := client.Get("name").Result()
	json, err := json.Marshal(User{Email: "abc@gmail.com", Password: "bca"})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(Ctx,"key1", json, 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get(Ctx, "key1").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val)
	}
}
