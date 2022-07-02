package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
//连接池
var db *sql.DB
var err error

func initDB(){
	dsn :="root:root@tcp(127.0.0.1:3306)/sql_test"
	db,err =sql.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("formatted error: %v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("mysql connection error: %v\n", err)
		return
	}
	fmt.Println("数据库连接成功！")
}
func queryOne(id int){
	//创建一个结构体对应查询的数据
	var user struct{
		id int
		name string
		age int
	}
		
	//用一个占位符代表id
	SQLstr :="SELECT id,name,age FROM users WHERE id=?"
	//Scan方法中注册了一个释放连接语句
	db.QueryRow(SQLstr,id).Scan(&user.id,&user.name,&user.age)
	

	db.SetMaxOpenConns(10)//设置连接池中最大连接数
	db.SetMaxIdleConns(5)//设置连接池中最大空闲连接数

	fmt.Printf("id=%d,name=%s,age=%d",user.id,user.name,user.age)

}
func queryMore(id int){
	SQLstr := "SELECT id,name,age FROM users WHERE id >?"

	//创建一个结构体对应查询的数据
	var user struct{
		id int
		name string
		age int
	}
	rows,err :=db.Query(SQLstr,id)
	if err != nil{
		fmt.Printf("error querying more failed! err: %v\n", err)
		return
	}
	//释放连接  请求多行数据时 Scan方法中没有释放连接
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&user.id,&user.name,&user.age)
		if err != nil {
			fmt.Printf("error scanning more failed! err: %v\n", err)
			return
		}
		fmt.Printf("id: %d name: %s age: %d\n", user.id, user.name,user.age)
	}
}
func main(){
	initDB()
	queryOne(2)
	queryMore(1)

}