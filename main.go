package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	for {

		fmt.Println("Go Redis Tutorial")

		//rediss: //:${module.redis-elasticache.auth_token}@${module.redis-elasticache.primary_endpoint_address}:6379
		address := "rediss://dummyuser:" + os.Getenv("REDIS_AUTH_TOKEN") + "@" + os.Getenv("REDIS_PRIMARY_ADDRESS") + ":6379"
		opt, err := redis.ParseURL(address)
		if err != nil {
			panic(err)
		}
		client := redis.NewClient(opt)
		t := time.Now()

		err = client.Set("date", t.String(), 0).Err()
		if err != nil {
			panic(err)
		}

		getValue, err := client.Get("date").Result()
		if err != nil {
			panic(err)
		}

		fmt.Println("Date set in elasticache:", getValue)
		time.Sleep(time.Minute)
	}

}
