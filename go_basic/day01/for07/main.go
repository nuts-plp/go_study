package main

import "fmt"

func main() {
	//for循环    go语言中循环只有for这一种

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	//变种1
	a := 0
	for ; a < 10; a++ {
		fmt.Printf("%d\n", a)
	}

	//变种2

	b := 0
	for b < 10 {
		fmt.Printf("%d\n", b)
		b++
	}

	//变种三    死循环
	// for{
	// 	fmt.Println("我好想你！潘丽萍")
	// }

	for i := 0; i < 10; i++ {
		if i == 5 { //当i==5时跳出循环
			break
		}
		fmt.Printf("%d\n", i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 { //当i==5时跳过此次循环
			continue
		}
		fmt.Printf("%d\n", i)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 3 && j == 5 {
				goto XX //当条件满足时，跳出循环到XX标签处
			}
			fmt.Printf("%d->%d\n", i, j)

		}
	}
XX: //Label标签
	fmt.Print("over")
}
