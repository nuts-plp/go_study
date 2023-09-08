package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Stu struct {
	Id   int `gorm:"primaryKey"`
	Name string
	Age  int
}
type Em struct {
	Id    int    `gorm:"autoIncrement;primaryKey"`
	Email string `gorm:"column:email"`
	StuId int    `gorm:"column:stu_id"`
}

func init() {
	//  通过一个现有的连接来初始化 *grom.DB
	dsn := "root:950629@tcp(47.92.232.226)/test?charset=utf8mb4&parseTime=True&loc=Local"
	coon, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: coon,
	}), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db
	//// 创建一个连接池来维护连接
	//sqlDB, err := DB.DB()
	//if err != nil {
	//	return
	//}
	//sqlDB.SetConnMaxLifetime(time.Hour)
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(100)
	//sqlDB.SetConnMaxIdleTime(time.Hour * 24)
}

// SQLCreate
// @Summary 创建数据记录
// @Description 通过结构体数据创建数据到数据库里
func SQLCreate() {
	//创建单条记录
	stu := Stu{
		6,
		"潘丽萍",
		18,
	}
	tx := DB.Table("stu").Create(&stu)
	fmt.Println("打印变动的row数：", tx)

	//创建多条记录
	stus := []*Stu{
		&Stu{7, "周周", 19},
		&Stu{8, "潘潘", 20},
	}
	rows := DB.Create(stus)
	fmt.Println("打印影响到的row：", rows)

	//根据单个map创建记录
	DB := DB.Table("stu").Create(map[string]interface{}{
		"id":   9,
		"name": "樊樊",
		"age":  19,
	})

	fmt.Println(DB)

	//根据map组创建记录
	DB.Model(&Stu{}).Create([]map[string]interface{}{
		{"id": 10, "name": "丽丽", "age": 18},
		{"id": 11, "name": "小丽", "age": 17},
		{"id": 12, "name": "君君", "age": 19},
		{"id": 13, "name": "笑话", "age": 15},
	})
}

// SQLSelect
// @Summary 查询数据记录
// @Description 查询数据记录并将其存储到结构体里
func SQLSelect() {

	// @单条查询

	// 实例化一个存储数据的对象
	stu1 := Stu{}

	//降序排序，查询第一个
	DB.Table("stu").First(&stu1)
	fmt.Println(stu1)

	//不排序，查询默认的第一个
	stu2 := Stu{}
	DB.Table("stu").Order("age").Take(&stu2)
	fmt.Println(stu2)

	//倒序查询第一个
	stu3 := Stu{}
	DB.Table("stu").Last(&stu3)
	fmt.Println(stu3)

	//根据主键值查询一个对象    注意：只能根据主键值
	stu4 := Stu{Id: 8}
	DB.Table("stu").Find(&stu4)
	fmt.Println(stu4)

	//// 实例化一个存储对象的map
	//result := map[string]interface{}{}
	//DB.Model(&Stu{}).First(result) //以此类推take last
	//
	//// @多条记录查询
	//stus := []*Stu{}
	//DB.Model(&Stu{}).Limit(2).Find(stus) //同理还有 where or offset select等
	//fmt.Println(stus)

	// @多条查询
	//根据主键值，也可不设置主键值，查询所有
	stus := []Stu{}
	tx := DB.Table("stu").Find(&stus, []int{7, 8, 9})
	fmt.Println(stus, tx.RowsAffected, tx.Error)

	//根据非主键值查询   string
	stu5 := Stu{}
	DB.Table("stu").Where("name=?", "潘潘").Find(&stu5)
	fmt.Println(stu5)

	//查询不等于     <>    select * from stu where name <> "周周";
	stu6 := []Stu{}
	DB.Table("stu").Where("name <> ?", "周周").Find(&stu6)
	fmt.Println(stu6)

	//查询 IN   select * from stu where name in ("周周","潘潘");
	stu7 := []Stu{}
	DB.Table("stu").Where("name in ?", []string{"潘潘", "周周"}).Find(&stu7)
	//DB.Table("stu").Find(&stu7, "name in ?", []string{"潘潘", "周周"})
	fmt.Println(stu7, "||||")

	//查询  > < select * from stu where age >18 and name in ("周周","潘潘");
	stu8 := []Stu{}
	DB.Table("stu").Where("age > ? and name in ?", 18, []string{"周周", "潘潘"}).Find(&stu8)
	//DB.Table("stu").Find(&stu8, "age >? and name in?",18,[]string{"周周", "潘潘"})
	fmt.Println(stu8)

	// between and    select * from stu where age between 17 and 20
	stu9 := Stu{}
	DB.Table("stu").Where("age between ? and ?", 18, 19).Find(&stu9)
	fmt.Println(stu9)

	//  指定字段内容   select * from stu where name = "潘潘";   下面三种方式等效
	stu10 := []Stu{}
	//DB.Table("stu").Where(&Stu{Name: "潘潘"}).Find(&stu10)
	//DB.Table("stu").Where(map[string]interface{}{
	//	"name": "潘潘",
	//}).Find(&stu10)
	//DB.Table("stu").Find(&stu10, Stu{Name: "潘潘"})
	DB.Table("stu").Find(&stu10, map[string]interface{}{
		"name": "潘潘",
	})
	fmt.Println(stu10)

	// not     select * from stu where not name = "潘潘";
	stu11 := []Stu{}
	//DB.Table("stu").Not("name = ?", []string{"潘潘", "周周"}).Find(&stu11)
	DB.Table("stu").Not(map[string]interface{}{
		"name": []string{"潘潘", "周周"},
	}).Find(&stu11)
	//DB.Table("stu").Where("name <> ?",[]string{"潘潘","周周"}).Find(&stu11)
	fmt.Println(stu11)

	// limit
	stu12 := []Stu{}
	DB.Table("stu").Limit(1).Find(&stu12, "name = ?", "潘潘")
	fmt.Println(stu12, 12)

	//order 排序  select * from stu order by age asc;
	stu13 := []Stu{}
	DB.Table("stu").Order("age asc").Find(&stu13)
	fmt.Println(stu13, 13)

	// limit&offset  偏移量 select * from stu offset 4 limit 3 order by id asc;
	stu14 := []Stu{}
	DB.Table("stu").Order("id asc").Offset(4).Limit(3).Find(&stu14)
	fmt.Println(stu14, 14)

	// select 选择特定字段
	stu15 := []Stu{}
	//DB.Table("stu").Select("name", "age").Find(&stu15)
	DB.Table("stu").Select([]string{"name", "age"}).Find(&stu15)
	rows, err := DB.Table("stu").Select("COALESCE(age,?)", 15).Rows()
	defer rows.Close()
	for rows.Next() {

	}
	fmt.Println(stu15, 15, rows, err)

	// group by &having  分组
	//stu16 := []Stu{}
	//DB.Table("stu").Group("name").Having("count(*) > ?", 1).Find(&stu16)
	//fmt.Println(stu16, 16)

	// 将结果存入一个数据库中没有表对应的变量中
	var h int
	DB.Table("stu").Select("sum(age) as su").Scan(&h)
	fmt.Println(h)

	// distinct  去重
	stu17 := []Stu{}
	DB.Table("stu").Distinct("name").Find(&stu17)
	fmt.Println(stu17, 17)
	// like 模糊查询
	stu18 := []Stu{}
	DB.Table("stu").Where("name like ?", "潘_").Find(&stu18)
	fmt.Println(stu18, 18)

}

// SQLUpdate
// @Summary 更新数据库记录
// @Description 更新数据库记录
func SQLUpdate() {}

// SQLDelete
// @Summary 删除数据记录
// @Description 从数据库中删除数据库记录
func SQLDelete() {}
func main() {
	//SQLCreate()
	SQLSelect()
	DB.Create(&Em{})
}
