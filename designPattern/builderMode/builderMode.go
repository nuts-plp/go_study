package main

import "fmt"

type Item interface {
	getName() string
	getPrice() float64
	Packing
}
type Packing interface {
	pack()
}
type Wrpper struct {
}

func (w *Wrpper) pack() {
	fmt.Println("纸盒装")
}

type Bottle struct{}

func (b *Bottle) pack() {
	fmt.Println("瓶装")
}

type Burger interface {
	getName() string
}
