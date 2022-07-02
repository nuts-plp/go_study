package main

import "fmt"

func main() {
	//Math.MaxFloat32     Math.MaxFloat64
	a:=1.2344
	fmt.Printf("%T\n",a)//go语言小数默认为float64

	b:=float32(1.234)//强制用float32位的
	fmt.Printf("%T\n",b)

}