package main

import "fmt"
import "os"
import "bufio"
import "io/ioutil"

func writeFile() {
	//打开文件  返回文件对象和error
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0632)
	//判断打开操作是否出现错误
	if err != nil {
		fmt.Printf("[writeFile]open file failed! Error:%v\n", err)
		return
	}
	defer file.Close()
	//要写入的数据
	str := "潘丽萍！好久不见，我好想你啊！"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("小潘！太想你了\n") //直接写入字符串数据

}

//bufio写入数据
func writeBufIo() {
	file, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("[writeBufio]open file failed! Error:%v\n", err)
		return
	}
	defer file.Close()
	//创建一个缓存文件对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Dear pan,we have too long to miss each other!\n")
	}
	//从缓存写入磁盘
	writer.Flush()
}

//用ioutil写入数据
func writeIoUtil() {
	str := "Honey pan,i miss you so much!\n"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 06512) //直接写入文件
	if err != nil {
		fmt.Printf("[writeIoUtil]write str to file failed! Error:%v\n", err)
		return
	}

}
func main() {
	//写入文件
	writeFile()
	writeBufIo()
	writeIoUtil()

}
