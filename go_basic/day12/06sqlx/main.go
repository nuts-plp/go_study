package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "database/sql"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	err error
)

func initDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("error check connecting failed!err:%v\n", err)
		return
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	fmt.Println("database initialized successfully!")

}

type user struct {
	id   int
	name string
	age  int
}

func main() {
	initDB()
	//单条查询
	SQLstr1 := "SELECT id,name,age FROM users WHERE id = 1"
	var u user
	db.Get(&u, SQLstr1)
	fmt.Printf("u:%#v\n", u)

	//多条查询
	SQLstr2 := "SELECT id,name,age FROM users "
	var userList []user
	db.Select(&userList, SQLstr2)
	fmt.Printf("userList:%#v\n", userList)

}
