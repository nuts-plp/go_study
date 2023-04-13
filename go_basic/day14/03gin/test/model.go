package main

type Book struct {
	ID    int64  `db:"ID"`
	Title string `db:"Title"`
	Price int64  `db:"Price"`
}
