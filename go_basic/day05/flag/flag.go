package flag

import (
	"flag"
	"fmt"
)

func TestFlag() {
	var (
		name     string
		age      int
		province string
	)
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.StringVar(&province, "province", "贵州", "省份")
	//flag.String(province, "贵州", "省份")
	flag.Parse()
	fmt.Println(name, age, province)
	fmt.Println(flag.Args())
}
