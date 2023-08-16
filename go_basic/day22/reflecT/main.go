package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	SayHello()
	Run()
	see()
}
type Leg struct {
	num int
}

type Hero struct {
	Name string `json:name`
	age  int    `json:age`
	Leg
}

func (h *Hero) SayHello() {
	fmt.Println("hello! " + h.Name + ".....")
}

func (h *Hero) Run() {
	fmt.Println("I am " + h.Name + ",我跑得飞快！")

}
func (h *Hero) see() {
	fmt.Println("I am " + h.Name + ",我看到了外星人！")
}
func main() {
	h := &Hero{"潘丽萍", 18, Leg{7}}
	h2 := Hero{"周小林", 19, Leg{9}}
	//获取结构体指针的类型
	typeofh := reflect.TypeOf(h)
	//结构体自身   结构体指针类型  结构体指针类型的种类   指针执行类型的真是实例对象
	fmt.Println(h, typeofh, typeofh.Kind(), typeofh.Elem())
	//获取结构体字段
	for i := 0; i < typeofh.Elem().NumField(); i++ {
		filed := typeofh.Elem().Field(i)
		//打印结构体字段的名称   标签tag   字段类型  字段是否匿名   字段偏移量
		fmt.Println(filed.Name, filed.Tag, filed.Type, filed.Anonymous, filed.Offset, filed.PkgPath)
	}
	fmt.Println("-----------------------------------")
	typeofh2 := reflect.TypeOf(h2)
	fmt.Println(h2, typeofh2, typeofh2.Kind())
	for i := 0; i < typeofh2.NumField(); i++ {
		filed := typeofh.Elem().Field(i)
		fmt.Println(filed.Name, filed.Tag, filed.Type, filed.Anonymous, filed.Offset, filed.PkgPath)
	}

	fmt.Println("============================================")
	valueOfH := reflect.ValueOf(h)
	fmt.Println(valueOfH, valueOfH.CanAddr(), valueOfH.CanSet(), valueOfH.Kind(), valueOfH.Type(), valueOfH.Elem())
	for i := 0; i < valueOfH.Elem().NumField(); i++ {
		field := valueOfH.Elem().Field(i)
		fmt.Println(field, field.Kind(), field.Type(), field.CanSet(), field.CanAddr(), valueOfH.Interface())

	}
	fmt.Println("|||||||||||||||||||||||||||||||||||||||||")
	valueOfh2 := reflect.ValueOf(h2)
	fmt.Println(valueOfh2, valueOfh2.CanAddr(), valueOfh2.CanSet(), valueOfh2.Kind(), valueOfh2.Type())
	for i := 0; i < valueOfh2.NumField(); i++ {
		field := valueOfh2.Field(i)
		fmt.Println(field, field.Kind(), field.Type(), field.CanSet(), field.CanAddr(), valueOfH.Interface())

	}
	reflect.ValueOf(h).Elem().Field(0).Set(reflect.ValueOf("fsdf"))
	var person Person = &Hero{"afasf", 18, Leg{9}}
	fmt.Println(reflect.TypeOf(person).Method(0).Func.Call([]reflect.Value{reflect.ValueOf(person)}))
	fmt.Println(reflect.TypeOf(person).Method(0))
}
