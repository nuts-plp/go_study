package main

import "fmt"
func main(){
	//make([]T,size,cap)使用make()函数创造新的切片，T为切片的数据类型，size为切片的数据长度，cap为切片的容量
	//make([]T,size)当没有指定容量时，默认容量和长度相等
	//使用make()函数创建一个切片
	a := make([]int,5,7)
	fmt.Printf("%v\n",a)

	//切片的本质是对底层数组的封装
	b := make([]int,0,0)
	fmt.Printf("%v\n",b)
	fmt.Println(b)

	c := make([]int,1)
	fmt.Printf("slice(c)->len:%d、cap:%d",len(c),cap(c))

	


}