package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func render(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./render.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, "你好！")
	//渲染模板

}
func main() {
	http.HandleFunc("/render", render)
	http.ListenAndServe(":9000", nil)
}
