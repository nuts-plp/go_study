package main

import "os"
import "fmt"
import "io"
import "bufio"
import "ioutil"


//读文件
func readFile(){
	//文件的读取
	//打开文件
	file , err := os.Open("./main.go")

	if err != nil{
		fmt.Printf("open file failed! Error:%v\n", err)
		return 
	} 
	//关闭文件
	defer file.Close()
	//读文件
	var temp [128]byte
	for{
		n , err :=file.Read(temp[:])//n为此次读取的字节数
		if err == io.EOF{
			fmt.Printf("read file failed! Error:%v\n", err)
		return
		}
		fmt.Print("读取了%d个字节！\n",n)
		fmt.Println("string(temp[:n])")
		if n < 128{
			return
		}
	}
	
}


//用bufio读文件
func readFileBufIo(){
	//打开文件
	file , err := os.Open("./main.go")
	if err == io.EOF{
		fmt.Printf("open file failed! Error:%v\n", err)
		return
	}
	//关闭文件
	defer file.Close()
	//创建一个读取内容的对象
	reader := bufio.NewReader(file)
	for{
		line, err := reader.ReadString('\n')

		if err == io.EOF{
			return
		}
		if err != nil{
			fmt.Printf("read file failed! Error:%v\n",err)
			return
		}
		fmt.Print(line)
	} 
	
}

//读取文件用ioUtil   读取整个文件
func readFileIoUtil(){
	ret ,err := ioutil.ReadFile("./main.go")
	if err != nil{
		fmt.Printf("read file failed! Error:%v\n",err)
		return
	}
	fmt.Println(string(ret))

}
func main(){
		readFile()
		readFileBufIo()
		// readFileIoUtil()	
	

}