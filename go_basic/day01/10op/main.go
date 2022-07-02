package main

import "fmt"

func main() {
	//运算符

	//算术运算符
	a:=5
	b:=3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	
	a++   //在go语言中是单独的运算符  不能放在=的右侧赋值， 等同于 a = a + 1

	a--		//同上  等同于 a = a - 1

	//关系运算符   go是强类型语言   必须是同类型才能做比较
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a < b)

	//逻辑运算符
	age := 22
	//与
	if age > 18 && age < 60{
		fmt.Println("苦逼的上班族！")
	}else{
		fmt.Println("不用上班，真好")
	}
	//或
	if age < 18 || age > 60{
		fmt.Println("不用上班的我们！")
	}else{
		fmt.Println("哎！要上班")
	}
	//非
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)
	fmt.Println("-------------------------------------")

	//位运算符
	//针对的是二进制数据
	//5对应的二进制数为101
	//2对应的二进制数为 10
	fmt.Println(5 & 2)//&按位与   两位均为1才为1          二进制结果为0 
	fmt.Println(5 | 2)//|按位或   两位有一个为1就为1		二进制结果为111  即7
	fmt.Println(5 ^ 2)//^按位异或    两位不一样为1			二进制结果为111  即7
	fmt.Println(5 << 2)//<<按位左移           二进制为10100   即20
	fmt.Println(5 >> 2)//>>按位右移          二进制位1   


	fmt.Println("---------------------------------")

	//赋值运算符
	var x int
	x = 10
	x += 1	//=>	x = x + 1
	x -= 1	//=>	x = x -1
	x *= 3	//=>	x = x * 3
	x /= 2	//=>	x = x / 2
	x &= 1	//=>	x = x & 1
	x |= 1	//=>	x = x | 1
	x ^= 1	//=>	x = x ^ 1
	x <<= 1	//=> x = x << 1
	x >>= 1	//=> x = x >> 1
	fmt.Println(x)

}