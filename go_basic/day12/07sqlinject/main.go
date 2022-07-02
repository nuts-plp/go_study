package main

import (
	"fmt"
	
)
var(
	db *sqlx.DB
	err error
)
type user struct {
	Id int
	Name string
	Age int
}
func initDB(){
	dsn :="root:root@tcp(127.0.0.1:3306)/sql_test"
	db,err = sqlx.Connect("mysql",dsn)
	if err != nil {
		fmt.Printf("Error check !err: %v\n", err)
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxIdleConns(5)
	fmt.Println("connection initialized successfully!")

}

func sqlInject(name string){
	SQLstr:=fmt.Sprintf("Select * from users WHERE name ='%s'",name)
	fmt.Printf("SQL: %s\n",SQLstr)

	var u user
	err = db.QueryRow(SQLstr).Scan(&u.Id,&u.Name,&u.Age)
	if err != nil{
		fmt.Printf("exec failed! err:%v\n",err)
		return
	}
	fmt.Printf("user:%#v\n",u)
}
func main(){
	initDB()
	sqlInject("潘丽萍")
	sqlInject("' or 1=1#")
	

}
