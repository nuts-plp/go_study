package main
import "fmt"
func main(){
	//go语言无法直接声明一个二进制常量，需要进行转换
	
	a:=90//十进制数
	fmt.Printf("%b",a)//转换为二进制数
	fmt.Printf("%o",a)//转换为八进制数
	fmt.Printf("%x",a)//转换为十六进制数
	fmt.Println()

	b:=077//八进制数
	fmt.Printf("%b",b)//转换为二进制数
	fmt.Printf("%d",b)//转换为十进制数
	fmt.Printf("%x",b)//转换为十六进制数
	fmt.Println()
	
	c:=0x123abc//十六进制数
	fmt.Printf("%b",c)//转化为二进制数
	fmt.Printf("%o",c)//转化为八进制数
	fmt.Printf("%x",c)//转化为十六进制数
}