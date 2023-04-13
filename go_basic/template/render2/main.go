package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	ID   int
}

func render2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./render2.tmpl")
	if err != nil {
		fmt.Println(err)
	}

	//渲染模板
	u1 := User{
		"潘丽萍",
		1901333,
	}
	m1 := map[string]interface{}{
		"name": "周小林",
		"age":  29,
	}
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})

}
func main() {
	http.HandleFunc("/render2", render2)
	http.ListenAndServe(":9000", nil)

}
