package main
import "fmt"


func main(){
	a:=[...]int{1,2,3,4,5,6,7,8}

	//计算这个数组的和
	var b int  = 0
	for _,i :=range a{
		b += i
		fmt.Printf("值:   %v\n",b)
		
	}

	//输出这个数组中两个数和为8的数的索引

	for i,j := range a{
		for n,m := range a{
			if (j + m) == 8{
				fmt.Printf("%d+%d=8,索引分别是%d,%d\n",j,m,i,n)
			}
		}
	}
}