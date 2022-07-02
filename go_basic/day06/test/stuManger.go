package main
import "fmt"


//学生管理系统
//1、存储学生数据
//2、保留一个管理者
type student struct {
	name string
	id int64

}


//管理者
type stuManger struct {
	allStudent map[int64]student
}



//显示学生信息
func (m stuManger)showStudent(){
	//循环遍历allStudent
	for _,v := range m.allStudent{
		fmt.Printf("学号:%d、姓名:%s\n",v.id,v.name)
	}

}

//添加学生
func (m stuManger)addStudent(){
	var newStuID int64
	var newStuName string

	//获取学生的信息
	fmt.Print("请输入学生的学号:")
	fmt.Scan(&newStuID)
	fmt.Print("请输入学生的姓名:")
	fmt.Scan(&newStuName)
	//根据输入数据创建结构体对象
	var newStudent = student{
		id:newStuID,
		name:newStuName,
	}
	m.allStudent[newStudent.id] = newStudent
	fmt.Println("添加学生信息成功！")

}

//删除学生信息
func (m stuManger)removeStudent(){
	//1、获取言删除的学生的学号
	fmt.Print("请输入要删除学生信息的学号:")
	var stuId int64
	fmt.Scan(&stuId)
	//根据输入的学号寻找对应的学生
	_, ok := m.allStudent[stuId]

	//如果根据学号没找到人
	if !ok{
		fmt.Println("查无此人！")
		return
	}
	//找到对应学号的人
	delete(m.allStudent,stuId)
	fmt.Println("删除成功！")

}

//修改学生信息
func(m stuManger)modifyStudent(){

	//1、获取言删除的学生的学号
	fmt.Print("请输入要修改学生信息的学号:")
	var stuId int64
	fmt.Scan(&stuId)
	//根据输入的学号寻找对应的学生
	stuObj, ok := m.allStudent[stuId]

	//如果根据学号没找到人
	if !ok{
		fmt.Println("查无此人！")
		return
	}
	//找到对应学号的人
	fmt.Printf("你要修改的学生学号:%d、姓名:%s\n",stuObj.id,stuObj.name)
	fmt.Print("请输入要修改的姓名:")
	var stuName string
	fmt.Scan(&stuName)
	//更新学生的姓名
	stuObj.name = stuName
	m.allStudent[stuId] = stuObj
	
	fmt.Println("修改信息成功！")

}