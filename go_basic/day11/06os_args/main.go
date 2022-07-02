package main
import(
	"fmt"
	"os"	
)

func main(){
	//从终端命令行获取参数   go run main.go -u -n -b -v
	fmt.Printf("%T\n",os.Args)
	fmt.Printf("%#v\n",os.Args)
	fmt.Printf("%#v\n", os.Args[0])
	fmt.Printf("%#v\n", os.Args[1])
	fmt.Printf("%#v\n", os.Args[2])
}