package main

import "fmt"
func main(){
	//append()函数为切片添加元素


	a := make([]string ,0,0)//创建一个空切片

	a = append(a, "上海","武汉")//为空切片a添加了一个元素
	fmt.Println(a)

	b := append(a,"深圳")//向a添加一个新元素后，把新切片赋值给了一个新变量
	fmt.Println(b)
	fmt.Println(a)
	fmt.Printf("slice(a)->%p\nslice(b)->%p",a,b)//通过索引地址发现底层数组不是同一个数组
	fmt.Println()


	s :=[...]string{"李光辉","潘丽萍"}
	s1 := s[:]//根据数组s创建切片s1
	s2 := []string{"周小林","牛敏","于欢"}
  
	s3 := append(s1,s2...)
	fmt.Println(s3)


	c := make([]int,6,6)//创建一个长度和容量都为6的切片
	c = append(c,2)
	c2 :=append(c,5)
	fmt.Println(c)
	fmt.Println(c2)
	fmt.Printf("slice(c)->%p\nslice(c2)->%p",c,c2)

	//利用append()函数实现切片元素的删除
	x :=[]int{0,1,2,3,4,5,6,7,8,9}

	x1 :=append(x[:3],x[4:]...)//删除数据3  此时的切片[0,1,2,4,5,6,7,8,9,9]
	fmt.Println(x)
	fmt.Println(x1)
	fmt.Printf("slice(x)->%p\nslice(x1)->%p",x,x1)
	fmt.Println()

	x2 :=append(x[:9],x[9:]...)
	fmt.Println(x2)

}