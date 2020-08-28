package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var url string = "mongodb://foo:bar@localhost:27017"

// 连接
func mongoClient ()  {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err == nil {
		fmt.Println("newclient", "新建连接失败")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err == nil {
		fmt.Println("connect", "连接失败")
	}
}

func main()  {
	// 连接mongodb
	mongoClient()
}