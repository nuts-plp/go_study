package main

import (
	"fmt"
	"go_basic/dataStruct/tree/AVLtree"
)

func _compare(a, b interface{}) int {
	var newA, newB int
	var ok bool
	if newA, ok = a.(int); !ok {
		return -1
	}
	if newB, ok = b.(int); !ok {
		return -2
	}
	if newA > newB {
		return -1
	} else if newB > newA {
		return 1
	} else {
		return 0
	}
}

func main() {
	myavl, _ := AVLtree.NewAVLTree(3, _compare)
	myavl = myavl.Insert(2)
	myavl = myavl.Insert(1)
	myavl = myavl.Insert(4)
	myavl = myavl.Insert(5)
	myavl = myavl.Insert(6)
	myavl = myavl.Insert(7)
	myavl = myavl.Insert(15)
	myavl = myavl.Insert(26)
	myavl = myavl.Insert(17)
	fmt.Println(myavl.GetAll())
}
