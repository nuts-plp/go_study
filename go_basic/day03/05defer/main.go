package main

import "fmt"

func main(){

	//defer
	//一般用于面试
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	defer fmt.Println("5")
	fmt.Println("end")

// fmt.Println(f1())//5
// fmt.Println(f2())//6
// fmt.Println(f3())//5
// fmt.Println(f4())//5

}
//5
func f1()int{
	x := 5
	defer func() {
		x++
	}()
	return x
}
//6
func f2()(x int){
	defer func() {
		x++
	}()
	return 5
}
//5
func f3()(y int){
	x := 5
	defer func() {
		x++
	}()
	return x
}
//5
func f4()(x int){
	defer func(x int) {
		x++
	}(x)
	return 5
}

