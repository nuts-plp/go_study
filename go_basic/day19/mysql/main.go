package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	db *sql.DB
)

type student struct {
	id   int
	name string
	age  int
}
type person struct {
	name string
	age  int
}

func InitDB() (err error) {
	str := "root:950629@tcp(47.143.67.231:3306)/student"
	db, err = sql.Open("mysql", str)
	if err != nil {
		fmt.Println("str syntax wrong")
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("mysql connect failed")
		return err
	}
	db.SetMaxIdleConns(4)
	db.SetConnMaxIdleTime(10)
	db.SetConnMaxLifetime(time.Second)
	return
}
func queryOne(id int) {
	var p student
	sqlStr := "SELECT id,name,age FROM user WHERE id=?"
	_ = db.QueryRow(sqlStr, id).Scan(&p.id, &p.name, &p.age)
	fmt.Println(p.id, p.name, p.age)

}
func queryRows(id int) {
	var p student
	sqlStr := "SELECT id,name,age FROM user WHERE id>?"
	rows, _ := db.Query(sqlStr, id)
	for rows.Next() {
		_ = rows.Scan(&p.id, &p.name, &p.age)
		fmt.Println(p.id, p.name, p.age)
	}
}
func execSql(name string, age int) {
	sqlStr := "INSERT INTO user(name,age)values(?,?)"
	result, _ := db.Exec(sqlStr, name, age)
	id, _ := result.LastInsertId()
	n, _ := result.RowsAffected()
	fmt.Println(id, n)
}
func sqlInject(name string) {
	var p student
	sqlStr := fmt.Sprintf("SELECT * FROM user where name=%s", name)
	_ = db.QueryRow(sqlStr).Scan(&p.id, &p.name, &p.age)
	fmt.Println(p.id, p.name, p.age)
}
func sqlPrepare(name string, age int) {
	sqlStr := "INSERT INTO user (name,age) VALUES (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare sql syntax failed")
		return
	}
	defer stmt.Close()
	stmt.Exec(name, age)
}
func main() {
	InitDB()
	var s = []*person{
		{"潘丽萍", 19},
		{"周小林", 22},
		{"于欢", 23},
		{"牛敏", 23},
		{"樊雪怡", 22},
	}
	for _, v := range s {
		sqlPrepare(v.name, v.age)
	}
	queryOne(3)
	queryRows(3)
	sqlInject("周小林")
	execSql("李光辉", 22)
}
