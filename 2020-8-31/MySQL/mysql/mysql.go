package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库参数配置
var (
	db *sqlx.DB
	dns string = "root:123456@tcp(192.168.196.227:3306)/users"
)

// 数据库表结构
type user struct {
	id int
	name string
}

// 连接数据库
func initDB(driverName string, dataSourceName string) (err error) {
	db, err = sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		fmt.Println("connect DB failed", err)
		return err
	}
	return err
}

// 根据id查询单条数据
func queryId(id int) {
	sqlStr := "select id, name from users_tbl where id=?"
	var u user
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u.id, u.name)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 根据名字查询单条数据
func queryName(name string) {
	sqlStr := "select id, name from users_tbl where name=?"
	var u user
	rows, err := db.Query(sqlStr, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 延后关闭
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u.id, u.name)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 查询全部数据
func queryAll()  {
	sqlStr := "select * from users_tbl"
	var u user
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u.id, u.name)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 插入
func insertIdName(id int, name string) {
	sqlStr := "INSERT INTO users_tbl (id, name) VALUES (?, ?)"
	_, err := db.Exec(sqlStr, id, name)
	if err != nil {
		fmt.Println("inset failed", err)
		return
	}
}

// 根据id更新数据
func updateId(id int, name string) {
	sqlStr := "update users_tbl set name = ? where id = ?"
	_, err := db.Exec(sqlStr, name, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 根据name更新数据
func updateName(id int, name string) {
	sqlStr := "update users_tbl set id = ? where name = ?"
	_, err := db.Exec(sqlStr, name, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 根据id删除
func deleteId(id int) {
	sqlStr := "delete from users_tbl where id = ?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 根据name删除
func deleteName(name string) {
	sqlStr := "delete from users_tbl where name = ?"
	_, err := db.Exec(sqlStr, name)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	err := initDB("mysql", dns)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("mysql连接成功")
	}
	{
		fmt.Println("根据id单条查询")
		queryId(1)
		// 键值不区分大小写
		fmt.Println("根据name单条查询")
		queryName("G")
		fmt.Println("全部查询")
		queryAll()
	}
	{
		queryAll()
		fmt.Println("插入数据")
		insertIdName(14, "gaolle")
		queryAll()
	}
	{
		queryAll()
		fmt.Println("根据id更新数据")
		updateId(5, "g")
		fmt.Println("根据name更新数据")
		updateName(5, "g")
		queryAll()
	}
	{
		queryAll()
		fmt.Println("根据id删除")
		deleteId(1)
		fmt.Println("根据name删除")
		deleteName("Gaolle")
		queryAll()
	}
}
