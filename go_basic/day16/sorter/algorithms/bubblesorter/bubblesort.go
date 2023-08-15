package bubblesorter

/*
冒泡排序,通过遍历和交换顺序来达到重新排序的目的，这是一个升序的排序方法.

*/
func BubbleSort(values []int) {
	flag := true
	for i := 0; i < len(values); i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}

}
