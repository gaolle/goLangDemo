package main

import (
	"fmt"
	"sqlx"
)

func main() {
	// redis
	// err := redis.Set("key", "val", 0)
	// fmt.Println(err)
	// value, err := redis.Get("key")
	// fmt.Println(value, err)
	// err := redis.Del("k4", "k3")
	// fmt.Println(err)
	// oldValue, err := redis.GetSet("key", "value")
	// if err != nil {
	// 	fmt.Println("redis getset err", err)
	// } else {
	// 	fmt.Println("redis getset old value:", oldValue)
	// }
	// mysql
	// if err := sqlxdb.Insert(10, "x"); err != nil {
	// 	fmt.Println("sqlx insert:", err)
	// }
	// if err := sqlxdb.Update(10, "B"); err != nil {
	// 	fmt.Println("sqlx update:", err)
	// }
	if err := sqlx.Delete(9); err != nil {
		fmt.Println("sqlx delete:", err)
	}
}