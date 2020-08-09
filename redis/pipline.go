package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.40.160:7001", "192.168.40.160:7002", "192.168.40.160:7003"},
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	fields := []string{"key1", "key2"}

	pipe := client.Pipeline()
	pipe.HMGet("yo", "key1", "key2")
	pipe.HMGet("yo1", "key1", "key2")
	re, err := pipe.Exec()

	fmt.Println(re, err)
	mp := make(map[string]map[string]string)
	for _, r := range re {
		sc, ok := r.(*redis.SliceCmd)
		if !ok {
			log.Fatal("err")
		}
		data := sc.Val()
		mpp := make(map[string]string)
		for i, field := range fields {
			if data[i] != nil {
				mpp[field], _ = data[i].(string)
			}
		}
		//fmt.Println(sc.Args())
		mp[sc.Args()[1].(string)] = mpp
	}
	fmt.Println(mp)

	ch := make(chan int)
	test(ch)
}

func test(ch chan int) {
	for e:= range ch  {
		fmt.Println(e)
	}
}
