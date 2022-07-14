package main

import "fmt"

func main() {
	s := "asdfghjas"
	l := len(s)
	b := 0
	for i := 0; i < l-b; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				if b < j-i {
					b = b - i
				}
			}
		}
	}
	fmt.Println(b)
}
