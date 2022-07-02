package main

import "fmt"

func main() {
	switch n:=3; n {
		case 1:
			fmt.Println("大拇指")
		case 2:
			fmt.Println("食指")
		case 3:
			fmt.Println("中指")
		case 4:
			fmt.Println("无名指")
		default:
			fmt.Println("不知名指")
	}

	a:= 7
	switch a {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	default:
		fmt.Println("不知名指")
	}

	switch n:=4;n{
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
		fallthrough		//为了兼容c语言设计的
	default:
		fmt.Println("不知名指")
	}
}