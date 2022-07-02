package split_string

import "strings"
// import "fmt"

func SplitStr(s string, sep string) []string {
	var index int
	var ret []string
	// fmt.Println("1234")
	for {
		index = strings.Index(s, sep)
		ret = append(ret,s[:index])
		s = s[index+len(sep):]
		// fmt.Println("123")
		if index <=0{
			ret = append(ret,s[index+1:])
			break
		}
	}
	
	// fmt.Println("12")
	return ret
}

// func main(){
// 	fmt.Println(Split("asfddssdf","s"))
// }