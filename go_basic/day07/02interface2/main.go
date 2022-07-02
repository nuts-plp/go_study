package main

import "fmt"

type eaterr interface{
	eat()
}

type drinkerr interface{
	drink()
}

type dog struct{
	name string
}

func(d *dog)eat(){
	fmt.Printf("我的名字叫%s、我是一只狗、我喜欢吃肉\n",d.name)
}

func (d *dog)drink(){
	fmt.Printf("我的名字叫%s、我是一只狗、我不喜欢喝水\n",d.name)
}

type cat struct{
	name string
}

func (c *cat)eat(){
	fmt.Printf("我的名字叫%s、我是一只猫、我喜欢吃鱼\n",c.name)
}

func (c *cat)drink(){
	fmt.Printf("我的名字叫%s、我是一只猫、我喜欢和可乐\n",c.name)
}


type flower struct{
	name string
}

func (f *flower)eat(){
	fmt.Printf("我是一朵花、我的名字叫%s、我喜欢阳光\n",f.name,)
}

func eater(e eaterr){
	e.eat()
}
func drinker(d drinkerr){
	d.drink()
}

func main(){
	d := dog{
		"旺财",
	}
	c := cat{
		"花斑猫",
	}
	f := flower{
		"太阳花",
	}
	eater(&f)
	eater(&c)
	fmt.Println("----------------------------")
	drinker(&d)
	drinker(&c)
}