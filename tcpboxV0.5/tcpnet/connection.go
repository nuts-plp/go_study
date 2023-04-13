package tcpnet

import (
	"errors"
	"fmt"
	"io"
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

	//告知当前链接已经退出/停止的channel
	ExitChan chan bool

	//消息的管理MsgID和对应的处理业务API关系
	MsgHandle tcpiface.IMsgHandle
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler tcpiface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		MsgHandle: msgHandler,
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
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if nil != err {
		//	fmt.Println("receive buf err", err)
		//	continue
		//}

		//创建一个装包拆包对象
		dp := NewDataPack()

		//读取客户端的msg head 二进制流8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); nil != err {
			fmt.Println("Read msg error:", err)
			break
		}
		//拆包， 得到msg的id和len 放入到msg中
		msg, err := dp.Unpack(headData)
		if nil != err {
			fmt.Println("unpack error:", err)
			break
		}

		//拆包，得到dataLen 再次读取Data，放在msg.Data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); nil != err {
				fmt.Println("read msg data error:", err)
				break
			}

		}
		req := Request{
			conn: c,
			msg:  msg,
		}

		//从路由中找到之前注册绑定的Conn对应的router调用
		//根据绑定好的MsgID找到对应处理api业务执行
		go c.MsgHandle.DoMsgHandle(&req)
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
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}
	dp := NewDataPack()
	//将data进行封包
	binaryMsg, err := dp.Pack(NewMsgPackage(msgID, data))

	if nil != err {
		fmt.Println("Pack error msg id :", msgID)
		return errors.New("Pack error msg")
	}

	//将数据发送到客户端
	if c.Conn.Write(binaryMsg); nil != err {
		fmt.Println("Write msg ID", msgID, "error:", err)
		return errors.New("conn Write error")
	}
	return nil
}
