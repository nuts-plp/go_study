package main
import(
	"fmt"
	"flag"
)
func main(){
	//标志位flag库  
	name := flag.String("name", "李光辉","请输入姓名")//三个参数分别为参数名、默认值、用法提示  返回值是指针
	age := flag.Int("age",0,"请输入年龄")
	sex := flag.String("sex", "male","请输入性别")
	time := flag.String("time", "00年","请输入时间")
	//flag.typeVar()
	var email string
	flag.StringVar(&email,"email","sncot123@aliyun.com","请输入邮箱")
	flag.Parse()//先解析再使用
	fmt.Printf("%T\n",*name)
	fmt.Printf("name:%v\n",*name)
	fmt.Printf("age:%v\n",*age)
	fmt.Printf("sex:%v\n",*sex)
	fmt.Printf("time:%v\n",*time)

	
	fmt.Printf("email:%v\n",email)


	fmt.Println(flag.Args()) //返回命令行参数后的其他参数 以[]string的形式
	fmt.Println(flag.NArg()) //返回命令行参数后其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数


}