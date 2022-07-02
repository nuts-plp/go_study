package main
 import "reflect"
 import "fmt"

 //使用typeof
func reflectTypeOf(x interface{}){
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n",v)
}
//使用valueof
func reflectValueOf(x interface{}){
	v := reflect.ValueOf(x)
	fmt.Printf("value:%v\n",v)
}
//使用kind
// func reflectKind(x interface{}){
// 	v := reflect.kind(x)
// }

func main(){
	v := "你好啊"
	reflectTypeOf(v)
	reflectValueOf(v)
	fmt.Printf("king:%v\n",reflect.TypeOf(v).Kind())

	var a float64 = 89
	reflectTypeOf(a)
	reflectValueOf(a)
	fmt.Printf("kind:%v\n",reflect.ValueOf(a).Kind())	

	
	var person struct {
		name string
		age int
	}
	person.name = "John"
	person.age = 90
	fmt.Printf("type:%v\n",reflect.TypeOf(person))
	fmt.Printf("type:%v\n",reflect.ValueOf(person))
	fmt.Printf("kind:%v\n",reflect.TypeOf(person).Kind())
	fmt.Println(person)
	fmt.Println("--------------------------------")

	var a1 *int
	fmt.Printf("a1:type(%v)kind(%v)\n",reflect.TypeOf(a1),reflect.TypeOf(a1).Kind())
	var a2 rune
	fmt.Printf("a1:type(%v)kind(%v)\n",reflect.TypeOf(a2),reflect.TypeOf(a2).Kind())
	var a3 map[int]string
	fmt.Printf("a1:type(%v)kind(%v)\n",reflect.TypeOf(a3),reflect.TypeOf(a3).Kind())
	var a4 bool
	fmt.Printf("a1:type(%v)kind(%v)\n",reflect.TypeOf(a4),reflect.TypeOf(a4).Kind())
	fmt.Printf("a1:type(%v)kind(%v)\n",reflect.TypeOf(main),reflect.TypeOf(main).Kind())

	b := 100
	reflectSetValue(b)

}

//使用elem()方法获取指针对应的值
func reflectSetValue(a interface{}){
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Int64{
		v.Elem().SetInt(400)
	fmt.Println(v)
	}
	
}