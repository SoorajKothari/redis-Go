package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx:= context.Background()
	client:= redis.NewClient(
		&redis.Options{
			Addr: "localhost:7000",
		},
	)
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Fail to ping redis")
		return;
	}
	fmt.Println("Connected to redis", pong)

	res, err := client.Set(ctx, "Name", "Sooraj", time.Minute).Result()
	if err != nil {
		fmt.Println("Fail to set key in redis")
		return;
	}
	fmt.Println("Set key success", res)

	value, err := client.Get(ctx, "Name").Result()
	if err != nil {
		fmt.Println("Fail to get key from redis")
		return;	
	}
	fmt.Println("Get returns ", value)

	val, err := client.Del(ctx, "Name").Result()
	if err != nil {
		fmt.Println("Fail to delete key from redis")
		return;	
	}
	fmt.Println("Deleted", val)
}
