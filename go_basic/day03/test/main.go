package main

import "fmt"


//小练习 分发金币
	//你有五十枚金币，需要分配给以下几个人：Mathew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabth
	//分配规则如下
	//a、名字中每包含一个e或E分一枚金币
	//b、名字中每包含一个i或I分二枚金币
	//c、名字中每包含一个o或O分三枚金币
	//d、名字中每包含一个u或U分四枚金币
	//写一个程序计算每个用户分到多少金币，剩下多少金币

//创建全局变量
var (
	coins = 50
	names = []string{"Mathew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabth"}
	distribution = make(map[string]int,len(names))	
)


//分发金币的方法
func dispatchCoins()(left int){
	//1、遍历每个名字
	for _,name := range names{
		
	//2、遍历每个名字中的字符

		for _,c := range name{

	//3、做字符匹配
			switch c {
				case 'e','E':
					distribution[name]++
	//4、分发金币
					coins--
				case 'i','I':
					distribution[name]+=2
					coins-=2
				case 'o','O':
					distribution[name]+=3
					coins-=3
				case 'u','U':
					distribution[name]+=4
					coins-=4
			}
		}
	}
	
	left = coins
	return
	
}
func main(){

	leftCoins := dispatchCoins()
	fmt.Printf("剩余金币:%d\n",leftCoins)
	for key,value := range distribution{
		fmt.Printf("%s:%d\n",key,value)
	}
}