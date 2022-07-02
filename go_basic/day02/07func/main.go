package main

import "fmt"

//定义函数
func f1(x int,y int)(ret int){
	return x + y
}

//匿名返回值
func f2(x int,y int)int{

	return x + y
}

//隐式返回
func f3(x int,y int)(ret int){
	ret = x + y
	return
}

//当参数类型相同时可以省略，前面的参数数据类型
func f4(x,y int)int{
	return x + y
}

//返回多个数据时
func f5(x,y int)(int,int){
	return x + y,x*y
}

//无返回的函数
func f6(x,y string){
	fmt.Printf("%s!%s",x,y)
}

//无返回值无参数的函数
func f7(){
	fmt.Println("我想你了！潘丽萍")
}

//可变长参数
func f8(x string,y ...int){
	fmt.Println(x)
	fmt.Println(y)//y是一个切片
}

//匿名内部函数
func f9(s string){

	vv := func(s string){

		fmt.Println(s," 我想你了")	
			
	}

	vv(s)
}

//只使用一次的匿名函数
func f10(x,y int){
	func(x,y int){
		fmt.Println(x,y,"what is this?")
	}(x,y)//直接调用
}

func main(){
	c :=f1(1,2)
	fmt.Println(c)
	d :=f2(6,7)
	fmt.Println(d)
	e :=f3(9,7)
	fmt.Println(e)
	b :=f4(3,2)
	fmt.Println(b)
	s,a :=f5(5,6)
	fmt.Println(a,s)
	f6("我爱你！","潘丽萍")
	f7()
	f8("你好",1,2,3,4,5)
	f9("潘丽萍")
	fmt.Println("--------------------------------------------")
	fmt.Printf("无参函数fu()->%T",fu)
	fmt.Println()
	fmt.Printf("有参函数fu1()->%T",fu1)
	fmt.Println()

	f10(10,100)



	//使用无参函数作为参数的函数
	fun1(fu)

	//使用有参函数作为参数的函数
	fun2(fu1)

	//使用无参函数作为返回值的函数
	df := fun3()
	df()
	fmt.Printf("函数fu1()->%T",fun3)
	fmt.Println()


	//使用有参函数作为返回值的函数
	dff := fun4()
	dff("牛敏!")
	fmt.Printf("函数fu1()->%T",fun4)
	fmt.Println()

	//使用有参函数作为参数使用有参函数作为返回值得函数
	dfff := fun5(fu2)
	dfff("潘丽萍")
	fmt.Printf("函数fu1()->%T",fun5)
	fmt.Println()

}
//无参函数
func fu(){
	fmt.Println("小潘！我想你了！")
}




//有参函数
func fu1(s string){
	fmt.Println(s," 我想死你了")
}

//使用无参函数作为参数的函数
func fun1(f func()){
	
	f()
}

//使用有参函数作为参数的函数
func fun2(ff func(string)){
	s := "周小林"
	ff(s)
}

//使用无参函数作为返回值的函数
func fun3()(ff func()){
	fv := "于欢！好久不见了"
	ff = func(){
		fmt.Println(fv)
	}
	return ff
}

//使用有参数的函数作为返回值的函数
func fun4()(func(string)){
	
	return func(s string){

		fmt.Println(s,"你好讨厌，我好喜欢")
	}

}

//使用有参函数作为参数使用有参函数作为返回值的函数

func fu2(s string){
	fmt.Println(s,"honey 你还好吗？")
}

func fun5(ff func(string))func(s string){
	ss := "wow! you are so beautiful!"
	ff(ss)
	return func(s string){
		fmt.Println(s,"when will we meet?")
	}
}