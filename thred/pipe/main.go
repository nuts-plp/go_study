package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

// IO管道
func IOPipe() {
	reader, writer := io.Pipe() //内存，管道
	go func() {
		output := make([]byte, 100)
		n, _ := reader.Read(output)
		fmt.Println("read:", n)
	}()
	input := make([]byte, 26)
	for i := 65; i < 90; i++ {
		input[i-65] = byte(i)
	}
	n, _ := writer.Write(input)
	fmt.Println("write:", n)
	time.Sleep(time.Second * 3)

}

// OS 系统管道
func OSPipe() {
	reader, writer, _ := os.Pipe() //系统，管道
	go func() {
		output := make([]byte, 100)
		n, _ := reader.Read(output)
		fmt.Println("read:", n)
	}()
	input := make([]byte, 16)
	for i := 65; i < 80; i++ {
		input[i-65] = byte(i)
	}
	n, _ := writer.Write(input)
	fmt.Println("write:", n)
	time.Sleep(time.Second * 3)

}

//匿名管道
func CMDPipe() {
	cmd1 := exec.Command("tasklist")
	var output bytes.Buffer //输出
	cmd1.Stdout = &output   //设置输出
	cmd1.Start()
	cmd1.Wait()
	fmt.Printf("%s    %v", output.Bytes(), output.Bytes())
}
func main() {
	//IOPipe()
	//OSPipe()
	CMDPipe()
}
