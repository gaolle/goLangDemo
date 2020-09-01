package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库参数配置
var (
	db *sqlx.DB
	dns string = "root:password@tcp(127.0.0.1:3306)/users"
)

// 数据库表结构
type User struct {
	Id int
	Name string
}

// 连接数据库
func init() () {
	var err error
	db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		panic(err)
		return
	}
	if db != nil {
		fmt.Println("connect DB")
	}
}

// 根据id查询单条数据
func selectById(id int) (User, error) {
	sqlStr := "select id, name from users_tbl where id=?"
	var u User
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Print("select id query: ", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Name)
		if err != nil {
			fmt.Println("select id scan: ", err)
		}
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("select id rows:", err)
	}
	return u, err
}

// 根据名字查询单条数据
func selectByName(name string) (User, error) {
	sqlStr := "select id, name from users_tbl where name=?"
	var u User
	rows, err := db.Query(sqlStr, name)
	if err != nil {
		fmt.Println("select name query: ", err)
	}
	// 延后关闭
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Name)
		if err != nil {
			fmt.Println("select name scan: ", err)
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("select name rows: ", err)
	}
	return u, err
}

// 查询全部数据
func selectAll()  {
	sqlStr := "select * from users_tbl"
	var u User
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("select all query: ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Name)
		if err != nil {
			fmt.Println("select all scan: ", err)
			return
		}
		fmt.Println("ID:", u.Id, "Name:", u.Name)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("select all rows: ", err)
		return
	}
}

// 插入
func insertIdAndName(id int, name string) error {
	sqlStr := "INSERT INTO users_tbl (id, name) VALUES (?, ?)"
	_, err := db.Exec(sqlStr, id, name)
	if err != nil {
		fmt.Println("insert failed", err)
		return err
	}
	return err
}

// 根据id更新数据
func updateById(id int, name string) error {
	sqlStr := "update users_tbl set name = ? where id = ?"
	_, err := db.Exec(sqlStr, name, id)
	if err != nil {
		fmt.Println("update by id failed", err)
		return err
	}
	return err
}

// 根据name更新数据
func updateByName(id int, name string) error{
	sqlStr := "update users_tbl set id = ? where name = ?"
	_, err := db.Exec(sqlStr, name, id)
	if err != nil {
		fmt.Println("update by name failed", err)
		return err
	}
	return err
}

// 根据id删除
func deleteById(id int) error {
	sqlStr := "delete from users_tbl where id = ?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete by id failed", err)
		return err
	}
	return err
}

// 根据name删除
func deleteByName(name string) error {
	sqlStr := "delete from users_tbl where name = ?"
	_, err := db.Exec(sqlStr, name)
	if err != nil {
		fmt.Println("delete by name failed", err)
		return err
	}
	return err
}

func main() {
	// 根据id查询单条数据
	if u, err := selectById(4); err != nil {
		fmt.Println("select by id failed")
	} else {
		fmt.Println("select by id success", u)
	}
}
