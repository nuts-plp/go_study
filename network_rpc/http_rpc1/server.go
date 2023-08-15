package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtils struct {
}
type Args struct {
	num1 *int
	num2 *int
}
type Reply struct {
	num1 *int
	num2 *int
}

func (m *MathUtils) CalculateCircleArea(a float32, resp *float32) error {
	*resp = math.Pi * a * a
	return nil
}
func (m *MathUtils) Say(a string, b *string) error {
	*b = "hello," + a
	return nil
}

func main() {
	//初始化指针数据类型
	mathUtils := new(MathUtils)
	//调用net/rpc包将服务对象注册
	_ = rpc.Register(mathUtils)
	//通过函数把mathUtil中提供的服务注册到http协议上，方便调用这可以利用http的方式进行数据传递
	rpc.HandleHTTP()
	//在特定的端口进行监听
	listen, err := net.Listen("tcp", "localhost:10001")
	if nil != err {
		panic(err)
	}
	http.Serve(listen, nil)
}
