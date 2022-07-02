package split
import (
	"strings"
)

func Split(str ,sep string)[]string{
	var ret =make([]string,0,strings.Count(str,sep)+1)
	index := strings.Index(str, sep)
	for index >=0{
		ret = append(ret,str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)

	}
	ret = append(ret,str)
	return ret
}

//斐波数列
func Fib(n int)int{
	if n <2{
		return n
	}
	return Fib(n-1)+Fib(n-2)
}