package main

func counter(count int) func() int {
	count++
	return func() int {
		return count
	}
}
func main() {
	counter(1)
}
