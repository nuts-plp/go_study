package main

import (
	"fmt"
	"strconv"
)

func main(){

	x := string(98)
	fmt.Printf("%#v   %T\n",x,x)

	//解析字符串为int数字
	a :="59"
	aValue,err :=strconv.Atoi(a)
	if err != nil {
		fmt.Println("Atoi failed!")
		return
	}
	fmt.Printf("%#v  %T\n",aValue,aValue)


	//解析数字为字符串
	a1 := 78
	a1Value := strconv.Itoa(a1)
	fmt.Printf("%#v   %T\n",a1Value,a1Value)
	//解析字符串为指定数据类型
	//解析字符串为bool
	a2 :="t"
	a2Value,err :=strconv.ParseBool(a2)
	if err != nil {
		fmt.Println("ParseBool failed!")
		return
	}
	fmt.Printf("%#v   %T\n",a2Value,a2Value)

	//解析字符串为int
	a3 := "-90"
	a3Value,err := strconv.ParseInt(a3,10,32)
	if err != nil {
		fmt.Println("ParseInt failed!")
		return
	}
	fmt.Printf("%#v   %T\n",a3Value,a3Value)

	

	//解析字符串为uint
	a4 := "900"
	a4Value,err := strconv.ParseUint(a4,10,32)
	if err != nil {
		fmt.Println("ParseUint failed!")
		return
	}
	fmt.Printf("%v   %T\n",a4Value,a4Value)


	//解析字符串为float
	a5 := "-90.875"
	a5Value,err := strconv.ParseFloat(a5,32)
	if err != nil {
		fmt.Println("ParseFloat failed!")
		return
	}
	fmt.Printf("%#v   %T\n",a5Value,a5Value)
}