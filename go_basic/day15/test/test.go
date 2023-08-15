package main

import "fmt"

/*
	this package is just for to test struct.

	just for testing struct
*/
type header map[string]interface{}

//add a k_v
func (i header) add(key, value string) {
	i[key] = value

}

//create a struct
type integer int

//add a number
func (i *integer) add(v integer) {
	*i += v
}

//main func to test struct
func main() {
	var a header = make(map[string]interface{}, 5)
	d := a
	d.add("mylover", "潘丽萍")
	fmt.Println(d["mylover"])
	var b integer = 1
	c := b
	c.add(5)
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println()
}

//BUG(sncot):#1 this is a issue
