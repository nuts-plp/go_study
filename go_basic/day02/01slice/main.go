package main

import "fmt"
func main(){
	//切片是引用数据类型 ，底层依然是数组

	//切片的声明    只声明没有初始化的切片是一个nil切片  nil切片的长度和容量都是0
	var a [] int   //此时的a是一个nil切片
	fmt.Println(a == nil)//go语言中的nil等同于其它语言中的null
	fmt.Printf("slice(a)->len:%d、cap:%d\n",len(a),cap(a))

	//1、切片的初始化方式一
	a = []int{1,2,3,4,5,5,6,7}

	//创建一个数组
	var b [9] int = [9]int{1,2,3,4,5,6,7,8,9}

	fmt.Printf("slice(a)type:%T、array(b)type:%T\n", a,b)

	//2、切片的初始化方式二
	//切片的容量：从切片开始对应的底层数组的位置到数组的结尾大小
	//切片的长度：字面意思
	c := b[:3]//由数组得到切片
	fmt.Printf("slice(c)len:%d、cap:%d\n",len(c),cap(c))

	d := b[5:7]
	fmt.Printf("slice(d)len:%d、cap:%d\n",len(d),cap(d))

	e := b[7:]
	fmt.Printf("slice(e)len:%d、cap:%d\n",len(e),cap(e))

	f := b[:]
	fmt.Printf("slice(f)len:%d、cap:%d\n",len(f),cap(f))

	//切片再切片,底层数组依然是最初的数组
	l := f[3:7]
	fmt.Printf("slice(l)len:%d、cap:%d\n",len(l),cap(l))

	b[8] = 100
	fmt.Printf("slice(e)%v\nslice(f)%v",e,f)

	//创建一个string数据类型的切片
	g := [4]string{"上海","武汉","广东","珠江"}

	j := g[:2]
	fmt.Printf("slice(j):%v\n",j)
	fmt.Println("----------------------------------")

	//切片的遍历
	//方式一
	for _, h := range j {
		fmt.Println(h)
	}
	//方式二
	for i := 0;i < len(j);i++{
		fmt.Println(j[i])
	}

	s := make([]string,0,0)  //长度和容量都为1的空切片，但不是nil切片
	fmt.Printf("slice(s)->len:%d、cap:%d\n",len(s),cap(s))
	fmt.Println("----------------------------------------------------------")

	v := [5]int{1,2,3} //
	v1 := v[:]
	fmt.Println(v1)
	fmt.Printf("slice(v)->len:%d、cap:%d\n",len(v1),cap(v1))
	fmt.Println()
	fmt.Println("----------------------------------------------------------")


	m := make([]int,3,5)
	fmt.Println(m)
	fmt.Printf("slice(m)->%p",m)
	fmt.Println()
	m = append(m,1)
	m = append(m,2)
	fmt.Println(m)
	fmt.Printf("slice(m)->%p",m)
	fmt.Println()
	
	
	m = append(m,1)//容量大于长度的部分用改数组类型的默认值填充
	fmt.Println(m)
	fmt.Printf("slice(m)->%p",m)
	fmt.Println()
	m[3] = 100
	m[2] = 200
	fmt.Printf("slice(m)->%p",m)
	fmt.Println()
	fmt.Println(m)
	fmt.Printf("slice(m)->%p",m)
	fmt.Println()
	m =append(m,4)//此时切片元素已超过切片长度，更换了底层数组
	fmt.Printf("slice(n)->%p",m)
	fmt.Println()
	n :=append(m,5)
	fmt.Printf("slice(m)->%p",n)
}