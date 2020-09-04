package sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// MySQL数据库对象
	db *sqlx.DB
)

const (
	dns string = "root:123456@tcp(127.0.0.1:3306)/users"
)

// 连接数据库
func init() {
	var err error
	db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		panic(err)
	}
	if db == nil {
		fmt.Println("sqlx no connect")
	} else {
		fmt.Println("sqlx connect")
	}
}

// 插入数据
func Insert(id int, name string) error {
	// 定义mysql语句
	sqlStr := "INSERT INTO users_tbl (id, name) VALUES (?, ?)"
	// 通过Exec插入数据
	_, err := db.Exec(sqlStr, id, name)
	return err
}

// 修改数据
func Update(id int, name string) error {
	// 定义mysql语句
	sqlStr := "update users_tbl set id = ? where name = ?"
	// 通过Exec插入数据
	_, err := db.Exec(sqlStr, id, name)
	return err
}

// 删除数据
func Delete(id int) error {
	// 定义mysql语句
	sqlStr := "delete from users_tbl where id = ?"
	// 通过Exec插入数据
	_, err := db.Exec(sqlStr, id)
	return err
}