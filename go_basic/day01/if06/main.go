package main

import "fmt"

func main() {
	//条件判断

	a:=19
	if a > 18{
		fmt.Printf("你已成年！可以进入酒吧！")
	}else{
		fmt.Printf("未成年!回家洗洗睡吧")
	}

	age:=20

	if age > 35{
		fmt.Println("孩子多大了！")
	}else if age >18{
		fmt.Println("你结婚了没？")
	}else{
		fmt.Println("赶紧回家写作业！")
	}
}