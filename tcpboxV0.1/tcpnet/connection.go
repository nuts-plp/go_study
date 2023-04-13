package tcpnet

import (
	"fmt"
	"net"
	"tcpbox/tcpiface"
)

/*
链接接口的实现
*/

type Connection struct {

	//当球按连接的socket
	Conn *net.TCPConn

	//连接的ID
	ConnID uint32

	//当前链接的状态
	isClosed bool

	//当前链接所i绑定的处理业务方法的API
	handleAPI tcpiface.HandleFunc

	//告知当前链接已经退出/停止的channel
	ExitChan chan bool
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi tcpiface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callbackApi,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

//链接的读数据方法
func (c *Connection) StartReader() {

	fmt.Println("Reader goroutine is running...")
	defer fmt.Println("ConnID:", c.ConnID, "Reader is exit,RemoteAddr is ", c.RemoteAddr().String())
	defer c.Stop()

	//读取刻划断的数据到buf中，最大512字节
	for {
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if nil != err {
			fmt.Println("receive buf err", err)
			continue
		}

		//调用当前链接所绑定的HandleApi
		if err := c.handleAPI(c.Conn, buf, n); nil != err {
			fmt.Println("ConnID:", c.ConnID, "handle failed... err:", err)
			break
		}
	}
}

//启动链接，让当前的链接准备工作
func (c *Connection) Start() {

	fmt.Println(" Conn Start()... ConnID:", c.ConnID)
	//启动当前连接的读数据业务
	go c.StartReader()

	//TODO 启动当前链接写数据业务

}

//停止链接，结束当前链接的工作
func (c *Connection) Stop() {

	fmt.Println(" Conn Stop()... ConnID:", c.ConnID)

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}

	c.Conn.Close()

	//关闭信道
	close(c.ExitChan)

}

//获取当前链接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端的TCP状态 IP PORT
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
