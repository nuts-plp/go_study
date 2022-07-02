package main
// import "./sort"
import "fmt"
// import "time"
// import "./rand"
// import "sort"
func main(){
	//map

	//声明map
	var a map[string]int
	//map的初始化
	a = make(map[string]int,10)
	
	//添加元素
	a["潘丽萍"] = 19
	a["周小林"] = 19
	a["于欢"] = 21
	a["牛敏"] = 22
	a["樊雪怡"] = 22
	fmt.Println(a)

	fmt.Println(a["李光辉"])
	_,ok := a["李光辉"]
	if !ok{
		fmt.Println("没有李光辉这个人")
	}else{
		fmt.Println("潘丽萍！我好想你")
	}

	//delete()函数删除键值对
	delete(a,"李光辉")//如果不存在，无操作
	

	//map的遍历
	for i,ok := range a {
		fmt.Println(i,ok)
	}

	// //map的指定顺序遍历
	// rand.Seed(time.Now().UnixNano())//初始化随机种子

	// var scoreMap = make(map[string]int,200)

	// for i:=0;i < 100;i++{
	// 	key := fmt.Printf("stu%02d",i)
	// 	value := rand.Intn(100)
	// 	scoreMap[key] = value
	// }

	// //取出map中所有key值存入切片
	// var keys = make([]string,0,200)
	// for key := range scoreMap{
	// 	keys = append(keys,key)
	// }

	// //对切片进行排序
	// sort.Strings(keys)

	// //按照排序后的key遍历map
	// for _,key := range keys{
	// 	fmt.Println(key,scoreMap[key])
	// }

	var s = make([]map[int]string,2,3)//声明一个数据类型为map的切片

	s[0] = make(map[int]string,2)//初始化map

	s[1] = make(map[int]string,3)//初始化map
	s[0][0] = "潘丽萍"
	s[0][1] = "周小林"
	s[1][0] = "于欢"
	s[1][1] = "牛敏"
	fmt.Println(s)
	

	v :=make(map[string][]string,5)
	fmt.Printf("%T",v)
	fmt.Println()
	V :=make(map[string][]rune,5)
	fmt.Printf("%T",V)


}