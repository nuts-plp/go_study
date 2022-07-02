package main

import (
	"fmt"
	"os"
)
//学生管理系统函数版


//操作菜单
func (m stuManger)showMenu(){
	fmt.Println(`----welcome Student management Systems----------
				1、显示学生信息
				2、添加学生信息
				3、删除学生信息
				4、修改学生信息
				5、退出系统
	`)

}

//声明一个管理者
var manager stuManger

func main(){
	manager = stuManger{
		allStudent:make(map[int64]student),
	}
	for {
		manager.showMenu()
		fmt.Print("请输入你的选择:")
		var a int
		fmt.Scanln(&a)
		fmt.Printf("你输入的数字是:%d\n",a)
		switch a {
		case 1:
			manager.showStudent()
		case 2:
			manager.addStudent()
		case 3:
			manager.removeStudent()
		case 4:
			manager.modifyStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("滚~")
			
		}
	}

}