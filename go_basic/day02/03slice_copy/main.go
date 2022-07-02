package main

import "fmt"
func main(){

	//切片的复制
	s1 := [...]string{"李光辉","潘丽萍","周小林","于欢","牛敏"}//创建一个数组

	s2 := s1[:]   //根据数组创建一个数据相同的切片
	fmt.Printf("slice(s2)->len:%d、cap:%d\n",len(s2),cap(s2))
	fmt.Println("------------------------------------------")

	//实现复制
	s3 := s2
	fmt.Printf("slice(s2)->%p\nslice(s3)->%p\n",s2,s3)

	fmt.Printf("array(s1)->%p",s1)

	//利用copy()函数复制切片
	s4 := make([]string, len(s2))
	copy(s4,s2)
	fmt.Println()
	fmt.Printf("slice(s4)->%p",s4)
	fmt.Println()
	fmt.Printf("slice(s2)->%p",s2)
	fmt.Println()
	fmt.Println(s4)



}