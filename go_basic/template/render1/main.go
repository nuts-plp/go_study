package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	ID     int
	gender string
}

func render1(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板

	p, err := template.ParseFiles("./render1.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	//渲染模板
	user := User{
		Name:   "潘丽萍",
		ID:     19,
		gender: "女",
	}
	err = p.Execute(w, user)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	http.HandleFunc("/render1", render1)
	http.ListenAndServe(":9000", nil)
}
