package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func test(r http.ResponseWriter, rq *http.Request) {}
func main() {
	file := filepath.FromSlash("d:/jsahd/asf/asdf/sdfsa/fasd")
	fmt.Println("|", file)

}
