package main

import "fmt"
func main() {
	var (
		name string
		age int 
		sex string
	)

	//注意： 扫描输入时必须使用指针
	//从控制台获取输入
	fmt.Scan(&name,&age,&sex)
	fmt.Printf("打印扫描结果name: %s、age: %d、sex: %s\n",name,age,sex)

	//格式化扫描终端输入 根据参数指定格式读取空白符分隔的值传给参数
	fmt.Scanf("%s %d %s\n", &name,&age,&sex)
	fmt.Printf("打印扫描结果name: %s、age: %d、sex: %s\n",name,age,sex)

	//遇到换行时才停止扫描最后一个数据后必须有换行或者到达结束位置
	fmt.Scanln(&name, &age, &sex)
	fmt.Printf("打印扫描结果name: %s、age: %d、sex: %s\n",name,age,sex)
}