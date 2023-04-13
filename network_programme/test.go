package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("X-Custom-Header", "dfsad")
	w.WriteHeader(201)

	fmt.Println(r)
	user := &User{
		"潘丽萍",
		18,
	}
	marshal, _ := json.Marshal(user)
	w.Write(marshal)
}
func main() {
	http.HandleFunc("/write", write)
	_ = http.ListenAndServe(":8080", nil)
}
