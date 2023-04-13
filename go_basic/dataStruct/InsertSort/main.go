package main

func InsertSort(source []int, asc bool) []int {
	var temp int
	for i := 0; i < len(source)-1; i++ {
		temp = source[i+1]
		var flag = false
		for j := 0; j < i+1; j++ {
			if asc {
				if source[j] > temp {
					flag = true
					source[j], temp = temp, source[j]
				}
				if flag {

				}
			} else {
				if source[j] < source[j+1] {
					source[j], source[j+1] = source[j+1], source[j]
				}
			}

		}
	}
	return source
}

func main() {

}
