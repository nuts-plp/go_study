package main

import "fmt"

//fmt占位符
func main() {
	a:=100
	fmt.Printf("二进制：%b\n",a)
	fmt.Printf("八进制：%o\n",a)
	fmt.Printf("十进制：%d\n",a)
	fmt.Printf("十六进制(表示为a-f)：%x\n",a)
	fmt.Printf("十六进制(表示为A-F)：%X\n",a)
	fmt.Printf("数据类型：%T",a)

	b:="你好！潘丽萍"
	fmt.Printf("字符串:%s\n",b)
	fmt.Printf("字符串：%v\n",b)
	fmt.Printf("字符串：%#v\n",b)
	fmt.Printf("引用 字符串：%#v\n",b)
	fmt.Printf("数据类型：%T\n",b)
	c := 'G'
	fmt.Printf("字符类型：%c\n",c)
	d := 65.0
	fmt.Printf("浮点型：%f\n",d)
	fmt.Printf("Unicode格式：%U\n",c)
	fmt.Printf("%q\n",65)//输出ASICC码中65对应的字母'A'


	

}