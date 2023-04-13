package main

import (
	"fmt"
	"regexp"
)

func main() {

	compile := regexp.MustCompile("faksdjfskajdfj")
	fmt.Println(compile.LiteralPrefix())
}
