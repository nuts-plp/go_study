package main
 
import "fmt"
import "os"
import "io"

//实现在文件夹的任意部位插入内容
func fileAdd(srcName,destName ,txt string,whe int){
	//打开指定文件
	srcFile,err := os.Open(srcName)
	if err == io.EOF{
		fmt.Printf("open %s failed! ",srcName)
		return
	}
	defer srcFile.Close()
	var where [1]byte
	srcFile.Read(where[:])

	destFile,err :=os.OpenFile(destName,os.O_RDWR|os.O_CREATE,06555)
	if err != nil{
		fmt.Printf("open %s failed! Error: %v\n",destName,err)
		return
	}
	defer destFile.Close()

	destFile.Write(where[:])
	//移动光标到要写入的位置
	srcFile.Seek(1,0)
	var str []byte
	srcFile.Read(str)
	destFile.WriteString(txt)


}


func main(){

}