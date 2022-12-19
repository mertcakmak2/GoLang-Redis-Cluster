package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

func main() {
	fmt.Println("redis cluster")
	ctx := context.Background()

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"redis1:6379", "redis2:6379", "redis3:6379", "redis4:6379", "redis5:6379", "redis6:6379"},
	})

	val, err := client.Ping(ctx).Result()
	fmt.Println(val, err)

	err = client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err = client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	forever := make(chan bool)
	<-forever

}
