package main

import "fmt"

func main() {

	//函数的递归调用
	
	
	fmt.Println(f1(10))
	fmt.Println(f2(10))

}


//求数的阶乘
//5*4!
//4*3!...
func f1(i int )(ii int){
	//设置结束退出条件
	if i==0 {
		return 1
	}

	return i * f1(i-1)
}

//面试题
//走楼梯，每次可以走一阶或者两阶，n阶楼梯有多少种走法

func f2(i int )(ii int){
	if i == 1 {
		return 1
	}
	if i == 2{
		return 2
	}

	return f2(i-1)+f2(i-2)
}