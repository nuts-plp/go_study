package main

import "net/http"

type Sncot struct{}

func (s *Sncot) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("你好 潘丽萍"))
}
func main() {

}
