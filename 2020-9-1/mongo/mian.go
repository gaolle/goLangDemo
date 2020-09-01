package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string
	Age int
}

var (
	client *mongo.Client
	dbName string = "student"
	dataName string = "student"
	url string = "mongodb://localhost:27017"
)

// 连接数据库
func init() {
	var err error
	clientOptions := options.Client().ApplyURI(url)

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to mongodb")
}

// 获取数据集
func GetData(client *mongo.Client) *mongo.Collection {
	return client.Database(dbName).Collection(dataName)
}

// 插入数据
func InsertData(collection *mongo.Collection, s Student) error {
	_, err := collection.InsertOne(context.TODO(), s)
	return err
}

// 删除数据
func DeleteDats(collection *mongo.Collection, d bson.D) (*mongo.DeleteResult, error) {
	res, err := collection.DeleteOne(context.TODO(), d)
	return res, err
}

// 查找数据
func SearchData(collection *mongo.Collection, d bson.D) (Student, error){
	var res Student
	err := collection.FindOne(context.TODO(), d).Decode(&res)
	return res, err
}

// 更新数据
func UpdateData(collection *mongo.Collection, object, operation bson.D) (*mongo.UpdateResult, error) {
	res, err := collection.UpdateOne(context.TODO(), object, operation)
	return res, err
}

// 关闭连接
func Close(client *mongo.Client) error {
	return client.Disconnect(context.TODO())
}

func main() {
	// s := Student{Name: "d", Age: 100}
	// s2 := Student{Name: "r", Age: 99}

	object := bson.D{{"name", "d"}}
	operation := bson.D{{"$set", bson.D{{"age", 99},}},}

	collection := GetData(client)

	// if err := InsertData(collection, s); err != nil {
	// 	fmt.Println("insert data", err)
	// } else {
	// 	fmt.Println("insert data success")
	// }

	// if res, err := DeleteDats(collection, filter); err != nil {
	// 	fmt.Println("delete data:", err)
	// } else {
	// 	fmt.Println("delete data count:", res.DeletedCount)
	// }

	// if res, err := SearchData(collection, filter); err != nil {
	// 	fmt.Println("search data:", err)
	// } else {
	// 	fmt.Println("search data value:", res)
	// }

	if res, err := UpdateData(collection, object, operation); err != nil {
		fmt.Println("update data", err)
	} else {
		fmt.Println("update data count", res.MatchedCount)
	}

	// if err := Close(client); err != nil {
	// 	fmt.Println("client close", err)
	// } else {
	// 	fmt.Println("close client success")
	// }
}
