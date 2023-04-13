package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

var (
	db *sqlx.DB
)

type student struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/student"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		err = fmt.Errorf("connect to mysql failed\n")
		return err
	}
	return
}

func get() {
	var p student
	sqlStr := "select id,name,age FROM user where id = ?"
	_ = db.Get(&p, sqlStr, 1)
	fmt.Println(p.Id, p.Name, p.Age)
}

func selectSql() {
	var p []student
	sqlStr := "SELECT id,name,age FROM user where id>?"
	_ = db.Select(&p, sqlStr, 1)
	fmt.Println(p)
}

//事务
func transactionDemo() (err error) {
	//事务的开始
	tx, err := db.Beginx()
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Println("commit")
		}
	}()
	sqlStr1 := "UPDATE user set age=age+2 WHERE id=:id"
	result, err := db.NamedExec(sqlStr1, map[string]interface{}{
		"id": 2,
	})
	if err != nil {
		return err
	}
	n1, _ := result.RowsAffected()
	if n1 != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "UPDATE user SET age=age-2 WHERE id=:id"
	result, _ = db.NamedExec(sqlStr2, map[string]interface{}{
		"id": 3,
	})
	n2, _ := result.RowsAffected()
	if n2 != 1 {
		return errors.New("exec sqlStr2 failed")
	}
	return err
}

func nameExec() {
	//借助变量名与插入字段相同，一一对应
	result, _ := db.NamedExec("INSERT INTO user(name,age)VALUES(:name,:age)", map[string]interface{}{
		"name": "潘丽萍",
		"age":  19,
	})
	fmt.Println(result.RowsAffected())
}
func nameQuery() {
	rows, _ := db.NamedQuery("SELECT id,name,age FROM user WHERE id=:id", map[string]interface{}{"id": 3})
	for rows.Next() {
		var u student
		rows.StructScan(&u)
		//rows.MapScan()
		//rows.SliceScan()
		fmt.Println(u)
	}
}

//自行构造批量插入的语句
func BatchInsertUsers(users []student) {
	//存放（？，？）的slice
	valueStrings := make([]string, len(users))
	//存放values的slice
	valueArgs := make([]interface{}, len(users)*2)
	//遍历users准备相关数据
	for _, u := range users {
		//此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?,?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	//自行拼接具体要执行的语句
	stmt := fmt.Sprintf("INSERT INTO user (name,age)VALUES %s", strings.Join(valueStrings, ","))
	_, _ = db.Exec(stmt, valueArgs)

}

//如股票要使用sqlx.In()那么结构体就要实现value方法
func (s student) Value() (driver driver.Value, err error) {
	return []interface{}{s.Name, s.Age}, nil
}

//如果student类是实现了value接口，则通过编译，否则报错
var _Value = &student{}

//sqlx.In实现批量的插入
func BatchInsertUser2(users []interface{}) {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name,age) VALUES (?),(?),(?)",
		users..., //如果args实现了driver.Value()  sqlx.In会通过调用Value()来展开他
	)
	fmt.Println(query)
	fmt.Println(args)
	_, _ = db.Exec(query, args)
}

//使用NameExec实现批量的插入
func BatchInsertUser3(users []*student) {
	db.NamedExec("INSERT INTO user (name,age)VALUES (:name,:age)", users)
}

func queryByIds(ids []int) (users []student) {
	//如果想以传入id的顺序返回查询结果
	// 1、借助代码  2、使用msyql内置函数find_in_set()
	strIDS := make([]string, 0, len(ids))
	for _, v := range ids {
		strIDS = append(strIDS, strconv.Itoa(v))
	}
	query, args, _ := sqlx.In("SELECT * FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id,?)", ids, strings.Join(strIDS, ","))
	//sqlx.In返回带？的bindvar的查询语句，我们使用ReBind重新绑定他
	query = db.Rebind(query)
	fmt.Println(query)
	db.Select(&users, query, args...)
	return
}
func main() {
	initDB()
	//get()
	//selectSql()
	//nameExec()
	//nameQuery()
	//transactionDemo()
	//s1 := &student{
	//	Name: "x",
	//	Age:  18,
	//}
	//s2 := &student{
	//	Name: "xx",
	//	Age:  28,
	//}
	//s3 := &student{
	//	Name: "xxx",
	//	Age:  38,
	//}
	////u := []interface{}{s1, s2, s3}
	////BatchInsertUser2(u)
	//u := []*student{s1, s2, s3}
	//BatchInsertUser3(u)
	s := queryByIds([]int{7, 5, 3, 1})
	fmt.Println(s)
}
