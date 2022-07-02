package main

import "fmt"
import "os"
import "io"


//实现文件的copy
func copyFile(srcName,desName string)(writen int64,err error){
	srcFile,err := os.Open(srcName)
	if err != nil{
		fmt.Printf("open %s failed! Error: %v\n",srcName,err)
		return
	}
	defer srcFile.Close()
	desFile,err := os.OpenFile(desName,os.O_WRONLY|os.O_APPEND|os.O_CREATE,06555)
	if err != nil{
		fmt.Printf("open %s failed! Error: %v\n",desName,err)
		return
	}
	defer desFile.Close()
	return io.Copy(desFile,srcFile)
}
func main(){
	copyFile("./main.go","./xx.txt")

}