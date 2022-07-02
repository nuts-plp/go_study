package main


import "fmt"

func main() {
	//数组的定义
	//数组的长度也是数组的一部分
	//数组声明时必须声明数组的类型和长度，数据为该类型的默认值
	var a [5] int
	fmt.Printf("%T",a)

	//1、初始化方式一
	a = [5]int{1,2,3,4,5}
	fmt.Println(a)

	//2、初始化方式二
	b:=[3]int{1,2,3}
	fmt.Println(b)

	c:=[...]string{"广州","上海","武汉"}//根据初始值自动推断数组长度
	fmt.Println(c)

	//3、初始化方式三  根据索引来初始化

	d:=[5]int{1:34,2:56}
	fmt.Println(d)

	for i:=0;i<len(c);i++{
		fmt.Printf("%d->%s\n",i,c[i])
	}

	for i,v := range c{
		fmt.Printf("%d--->%s\n",i,v)
	}


	//多维数组
	var e [3][2] int
	e = [3][2]int{
		[2]int{1,2},
		[2]int{2,3},
		[2]int{3,4},
	}
	fmt.Println(e)


	//多维数组的遍历
	for _,v := range e{
		fmt.Println(v)
		for _,m := range v{
			fmt.Println(m)
		}
	}

}