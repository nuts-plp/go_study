package main

import "fmt"

func bubbleSort(source []int, asc bool) []int {
	var length = len(source)

	for i := 0; i <= length; i++ {
		for j := 0; j+1 < length-i; j++ {
			//如果要求升序排序
			if asc {
				if source[j] > source[j+1] {
					source[j], source[j+1] = source[j+1], source[j]
				}
			} else { //降序排序
				if source[j] < source[j+1] {
					source[j], source[j+1] = source[j+1], source[j]
				}
			}

		}
	}
	return source
}
func main() {
	a := []int{2, 4, 5, 2, 4, 3, 5, 6, 2, 1, 4}
	sort := bubbleSort(a, false)
	fmt.Println(sort)
}
