package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func InitDB() (err error) {
	db, err = sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/books")
	if err != nil {
		fmt.Println("database connected failed!")
		return
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	return nil

}

func InsertBook(title string, price int) (err error) {
	sqlStr := "Insert INTO book(title,price) values (?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("insert failed!")
		return
	}
	return nil
}
func SelectBooks() (bookList []*Book, err error) {
	sqlStr := "SELECT id,title,price FROM book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("query failed!")
		return
	}
	return
}
func DeleteBook(id int64) (err error) {
	sqlStr := "DELETE FROM book WHERE id=?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete failed!")
		return
	}
	return
}
