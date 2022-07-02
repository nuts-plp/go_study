package main

import (
	"fmt"
	"math/rand"
	"time"
)

//rand

func main(){
	//留意加种子与不加种子的区别
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().UnixMicro())
	for	i := 0; i < 10; i++ {
		fmt.Println(rand.Int())
	
	}


}



