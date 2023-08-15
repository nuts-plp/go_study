package main

import (
	"fmt"
	"reflect"
)

type animal struct {
	leg int
}
type dog struct {
	animal
	bark string
}
type pc interface {
	eat()
	run() string
}

func (d dog) eat() {
	fmt.Println("我会" + d.bark)
}
func (d dog) run() string {
	return fmt.Sprintf("我有", d.leg, "条腿!")
}
func test() {
	fmt.Println("fghjk")
}

func main() {
	a := &dog{animal{4}, "汪汪"}
	dogOfType := reflect.TypeOf(a)
	fmt.Printf("%v- %v- %v\n", dogOfType, dogOfType.Elem(), dogOfType.Elem().Kind())
	for i := 0; i < dogOfType.Elem().NumField(); i++ {
		do := dogOfType.Elem().Field(i)
		fmt.Printf("%v- %v \n", do, do.Type)
	}
	dogOfValue := reflect.ValueOf(a)
	fmt.Printf("%v- %v- %v\n", dogOfValue, dogOfValue.Type(), dogOfValue.Kind())
	fmt.Println(dogOfValue.Elem().NumField())
	for i := 0; i < dogOfValue.Elem().NumField(); i++ {
		do := dogOfValue.Elem().Field(i)
		fmt.Printf("%v- %v- %v- %v- %v- %v\n", do, do.Type(), do.Kind(), do.Addr(), do.CanAddr(), do.CanSet())
	}
	fmt.Println("--------==================-----------------------")
	b := dog{animal{4}, "喵喵"}
	catOfType := reflect.TypeOf(b)
	fmt.Printf("%v %v \n", catOfType, catOfType.Kind())
	for i := 0; i < catOfType.NumField(); i++ {
		cat := catOfType.Field(i)
		fmt.Printf("%v- %v\n", cat, cat.Type)
	}
	catOfValue := reflect.ValueOf(b)
	fmt.Printf("%v- %v- %v\n", catOfValue, catOfValue.Type(), catOfValue.Kind())
	fmt.Println(catOfValue.NumField())
	for i := 0; i < catOfValue.NumField(); i++ {
		cat := catOfValue.Field(i)
		fmt.Printf("%v- %v- %v- %v\n", cat, cat.Type(), cat.Kind(), cat.CanSet())
	}
	c := "小红"
	reflect.ValueOf(&c).Elem().Set(reflect.ValueOf("moumou"))
	fmt.Println(c)
	fmt.Println("||||||||||||||||||||||||||||||||||||||||||||||||||")

	method := reflect.ValueOf(test)
	method.Call(nil)

	var k pc = &dog{
		animal{4},
		"wangwang",
	}
	of := reflect.TypeOf(k)
	mathod, _ := of.MethodByName("run")
	mathod.Func.Call(nil)
}
