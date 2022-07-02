package main

import "fmt"
import "strings"

func main() {
	//字符串用双引号""
	a:="你好！"
	//字符用单引号''
	b:='a'
	fmt.Printf("%s%v",a,b)
	//一个字节 Byte 是8个比特（比特（二进制位））
	//一个比特是一个二进制位
	//一个utf-8编码的中文一般是3个字节
	
	c:=`
	
	枯藤老树昏鸦
		小桥流水人家
			古道西风瘦马
				夕阳西下
					断肠人在天涯

	`
	fmt.Println(c)

	path:="C:\\ProgramData\\Microsoft OneDrive\\setup"

	fmt.Println(path)
	d:=fmt.Sprintf("%s%v",a,b)  //拼接字符串，精品街好的字符串返回给d
	fmt.Println(d)

	ss:=a+"潘丽萍"	//字符串拼接
	fmt.Println(ss)
	fmt.Println(len(path))//返回字符串长度

	fmt.Println(strings.Split(path,"\\"))//以\分割path

	fmt.Println(strings.Contains(c,"鸦"))//判断是否包含

	fmt.Println(strings.HasPrefix(ss,"枯藤"))//判断字符串是否以“枯藤”开始

	fmt.Println(strings.HasSuffix(ss,"丽萍"))//判断字符串是否以“丽萍”结尾

	sc:=strings.Split(path,"\\")
	fmt.Println(strings.Join(sc,"-"))

	sv:="abcdcdfsae"
	fmt.Println(strings.Index(sv,"c"))//返回首次出现位置
	fmt.Println(strings.LastIndex(sv,"a"))//返回最后一次出现位置的索引

}