package main
import(
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
var db *sql.DB
var err error
//初始化连接
func initDB(){
	dsn:="root:root@tcp(127.0.0.1:3306)/sql_test"

	db,err = sql.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("error opening database failed! err: %v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("error connecting to database failed! err: %v\n", err)
		return
	}
	fmt.Println("database initialized successfully!")
}

//插入数据
func insert(id int,name string,age int){
	SQLstr :="INSERT INTO users (id,name,age) VALUES (?,?,?)"
	result,err :=db.Exec(SQLstr,id,name,age)
	if err != nil {
		fmt.Printf("error inserting into database failed! err: %v ",err)
		return
	}
	ID,err := result.LastInsertId()
	if err != nil {
		fmt.Printf("error get lastInsertId failed! err: %v\n",err)
		return
	}
	fmt.Printf("insert successfully! lastInsertId:%d\n",ID)
}

//更新数据
func update(age,id int){
	SQLstr := "UPDATE users SET age=? WHERE id=? "
	result,err :=db.Exec(SQLstr,age,id)
	if err != nil {
		fmt.Printf("error updating database failed! err: %v\n",err)
		return
	}
	rows,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("error check affected rows!err:%v\n",err)
		return
	}
	fmt.Printf("update successfully! affected rows:%d\n",rows)

}

//删除数据
func delete(id int){
	SQLstr:="DELETE FROM users WHERE id=?"
	result,err :=db.Exec(SQLstr,id)
	if err != nil {
		fmt.Printf("error deleting database failed! err: %v\n",err)
		return
	}
	rows,err:= result.RowsAffected()
	if err != nil {
		fmt.Printf("error check affected rows!err:%v\n",err)
		return
	}
	fmt.Printf("delete successfully! affected rows:%v\n",rows)

}

func main(){
	initDB()
	insert(6,"樊雪姨",22)
	update(23,6)
	delete(4)

}
