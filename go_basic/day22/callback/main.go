package main

import "fmt"

func test(name string) {
	fmt.Println("I am " + name)
}

func fun(name string, t func(str string)) {
	t(name)
}
func main() {
	fun("潘丽萍", func(str string) {
		fmt.Println("I am " + str)
	})
}
