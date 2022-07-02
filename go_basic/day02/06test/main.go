package main

import (
	"strings"
	"fmt"
	"unicode"
)

func main(){
	//作业
	//1、计算一个字符串里有多少汉字
	s := "hello! dear honey 潘丽萍 "
	a := wordsCount()
	fmt.Println(a)

	//2、计算一句话中单词出现的次数
	ss :="How do you do？"
	b := mapCount(ss)
	fmt.Println(b)

	//3、回问判断
	//字符串从左往右和从右往左读是一样的
	//例：   上海自来水来自海上
	//		 山西运煤车运煤西山
	//       黄山落叶松叶落山黄
	huiText("黄山叶落松叶落山黄")
}




//传入一个字符串，返回字符串中汉字的个数
func wordsCount(s string)int{
	count := 0
	for _, word := range s{
		if word == unicode.Is(unicode.Han,word){
			count++
		}
	}
	return count
}


//传入一个字符串，返回一个map
func mapCount(s string)(m map[string]int){
//1、创建一个用于存储的map
	m = make(map[string]int,10)
//2、把字符串以空格分隔，得到切片
	v := strings.Split(s," ")
//3、遍历切片，把词存入map
	for _, word := range v{
//4、判断map中是否有这个key，如果没有+1
		if _,ok := map[word];!ok{
			map[word] = 1
		}else{
//5、map中存在这个key，值++
			map[word]++
		}
	}
	return m
}

func huiText(s string){
//1、创建一个切片用于存储字符
	ss := make([]rune, len(s))
	for _, word := range s{
		ss = append(ss,word)
	}
//2、用for循环进行判断并打印
	for i := 0;i < len(ss);i++{
		if ss[i] != ss[len(ss)-1-i]{
			fmt.Println("该字符串不是回文！")
			return
		}
	}
	fmt.Println("该字符串是回文！")
}