package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	dsn :="root:root@tcp(127.0.0.1:3306)/shop"
	db,err :=sql.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("Error opening	err: %v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("database connection error: %v\n", err)
		return
	}
	fmt.Println("database connection established!")
}