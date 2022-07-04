package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func initDB() {

	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("format check failed! err:%c\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("connected to database failed! err:%v\n", err)
		return
	}
	fmt.Println("database connected successfully!")

}

//prapare 预处理
func prepare() {

	type user struct {
		id   int
		name string
		age  int
	}
	var u = map[int]user{
		8:  {8, "李四", 20},
		9:  {9, "王五", 21},
		7:  {7, "张三", 17},
		10: {10, "赵六", 29},
	}

	SQLstr := "INSERT INTO users (id,name,age) VALUES (?,?,?)"

	stmt, err := db.Prepare(SQLstr)
	if err != nil {
		fmt.Printf("prepare failed! err:%v\n", err)
		return
	}
	defer stmt.Close()
	for _, v := range u {
		stmt.Exec(v.id, v.name, v.age)
	}

}
func main() {
	initDB()
	prepare()
}
