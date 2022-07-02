package main

import "reflect"
import "fmt"

type person struct {
	Name string `json:"name`
	Age int	`json:"age"`
	Gender string `json:"gender"`
	Number string `json:"number"`
	Address string `json:"address"`
	City string `json:"city"`
	Email string `json:"email"`
}

func main(){
	a := person{
		"潘丽萍",
		19,
		"女",
		"1678908765",
		"贵州省",
		"黔东南",
		"768901398@qq.com",
	}
	t := reflect.TypeOf(a)
	//循环遍历结构体所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s  index:%d  type:%v  json tag:%v\n",field.Name,field.Index,field,field.Tag.Get("json"))
	}

	//通过字段名获取指定的结构体信息
	if nameField,ok := t.FieldByName("name");ok{
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n",nameField.Name,nameField.Index,nameField.Type,nameField.Tag.Get("json"))
	}
}