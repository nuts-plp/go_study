package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
var(
	db *sql.DB
	err error
)
//初始化连接
func initDB(){
	dsn:="root:root@tcp(127.0.0.1；3306)/sql_test"

	db,err = sql.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("error opening database: %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("error connecting to database: %v", err)
		return
	}
	fmt.Println("database connected	successfully!")
}

func transaction(){
	tx ,err:= db.Begin()
	if err != nil {
		fmt.Printf("error starting transaction: %v", err)
		return
	}
	SQLstr1:="UPDATE users SET age = age - 2 WHERE id = 2"
	SQLstr2:="UPDATE users SET age = age + 2 WHERE id = 3"

	_,err = db.Exec(SQLstr1)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql语句1执行失败！回滚")
		return
	}
	_,err = db.Exec(SQLstr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("sql语句2执行失败！回滚")
		return
	}
	tx.Commit()
	fmt.Println("事务处理结束！")

}
func main(){
	initDB()
	transaction()
}