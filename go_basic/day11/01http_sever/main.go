package main

import (
	"io"
	// "bufio"
	"net/http"

	// "ioutil"
	"fmt"
	"os"
)
func f1(responseWriter http.ResponseWriter, r *http.Request){
	// text,err := ioutil.ReadFile()
	// str := `<h1 style="color:green">你好啊！潘丽萍</h1>`
	

	responseWriter.Write(*OpenFile1())

}

//OpenFile 打开并读取文件
func OpenFile1()*[]byte{
	//1、打开文件
	file,err := os.Open("./test.html")
	if err != nil{
		fmt.Println("open file failed! err:", err)
		return nil
	}
	defer file.Close()
	var content []byte
	txt := make([]byte,128)
	//2、读取内容存入切片
	for{
		n, err := file.Read(txt)
		if err == io.EOF{
			fmt.Println("文件读完了！")
			break
		}
		if err != nil {
			fmt.Println("read file failed! err:",err)
			return nil
		}
		content = append(content,txt[:n]...)
	}
		
	
	return &content
}

//OpenFile2 bufio读取文件
// func OpenFile2() *[]byte{
// 	file,err := os.Open("./test.html")
// 	if err != nil{
// 		fmt.Println("open file failed! err:",err)
// 		return nil
// 	}
// 	newReader := bufio.NewReader(file)
// 	context := make([]byte,128)
// 	//2、循环读取文件内容加入到context切片中
// 	for{
// 		line,err :=newReader.ReadString('\n')
// 		if err == io.EOF{
// 			if len(line) > 0{
// 				context = append(context,[]byte(line)...)
// 			}else{
// 				break
// 			}
			
// 		}
// 		if err != nil{
// 			fmt.Println("read file failed! err:",err)
// 		}
// 	}
// 	return &context
// }

func f2(rw http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(r.Body)
	rw.Write([]byte("ok"))
}

func main(){
	//
	http.HandleFunc("/hello/",f1)
	http.HandleFunc("/qury/",f2)
	http.ListenAndServe("127.0.0.1:9999",nil)
	
}