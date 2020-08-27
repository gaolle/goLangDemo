package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var client *redis.Client

// 连接redis服务器
func redisConnect() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	if pong != "PONG" {
		log.Fatal("连接失败")
	}else {
		fmt.Println("连接成功")
	}
}

// string
// get
func get(key string) {
	res, err := client.Get(key).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get", res)
}

/*
*
*/
// set
func set(key, val string) {
	res, err := client.Set(key, val, 0).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("set", res)
}

// hash
// hset
func hset(hashTable, key, val string) {
	res, err := client.HSet(hashTable, key, val).Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hset", res)
}

// hget
func hget(hashTable, key string) {
	res, err := client.HGet(hashTable, key).Result()

	if err != nil {
		log.Fatal(err)
	}
	// 返回值为是否新建field 返回0表示未新建
	fmt.Println("hget", res)
}

func main() {
	redisConnect()
	// set 操作
	get("key")
	set("key", "value")
	// hash 操作
	hget("k", "f")
	hset("k", "f", "value")
}
