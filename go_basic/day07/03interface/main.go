package main

import "fmt"

type mover interface{
	move()
}

type car struct{

}

func (c car) move() {
	fmt.Println("开车了！速度10迈")
}
type dog struct{
	animal
}

type animal struct{

}

func (a *animal) move() {
	fmt.Println("run~速度15迈")
}
func main(){
	var m mover
	fmt.Println(m)

	c := &car{}
	m = c
	m.move()
	fmt.Println(m)
	c1 := car{}
	m = c1
	fmt.Println(m)

	a := &animal{}
	a1 := &animal{}
	//判断一下能否把父类的方法调用
	m = a
	fmt.Println(m)

	m = a1
	fmt.Println(m)

	d := dog{}
	d.move()


}