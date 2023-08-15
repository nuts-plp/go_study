package tcpnet

import (
	"fmt"
	"net"
	"tcpbox/tcpiface"
	"tcpbox/utils"
)

//IServer的接口实现，定义一个Server的服务器模块

type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听的ip
	IP string
	//服务武器监听的端口
	Port int
	//当前的server添加一个router，server注册的链接对应的处理业务
	MsgHandler tcpiface.IMsgHandle
}

func (s *Server) Start() {

	//在启动之初打印一下zinx的配置信息
	fmt.Printf("[Zinx] Name:%s\n", utils.GlobalObject.Name)
	fmt.Printf("[Zinx] Host:%s,Port:%d\n", utils.GlobalObject.Host, utils.GlobalObject.Port)
	fmt.Printf("[Zinx] Version:%s,MaxConn:%d,MaxPackageSize:%d\n", utils.GlobalObject.Version,
		+utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)

	//用一个goroution来处理

	go func() {
		//开启消息队列及worker工作池
		s.MsgHandler.StartWorkerPool()

		fmt.Printf("[Start] Server Listener at %s ,Port %d,is starting\n", s.IP, s.Port)
		//1、获取一个TCP的ADDR
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if nil != err {
			fmt.Println("resolve tcp addr err:", err)
			return
		}

		//2、监听TCP
		ip, err := net.ListenTCP(s.IPVersion, addr)
		if nil != err {
			fmt.Println("listen ip err:", err)
			return
		}
		//3、阻塞等待连接
		fmt.Println("Start Zinx Server successfully,Name: ", s.Name, ",Listening")
		//链接ID
		var cid uint32 = 0
		for {
			//循环等待客户端连接
			conn, err := ip.AcceptTCP()
			if nil != err {
				fmt.Println("accept err:", err)
				continue
			}
			defer conn.Close()

			//将处理新链接的业务方法与Conn绑定得到我们的链接模块
			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++
			go dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {
	//TODO 将一些服务器的资源、链接进行停止、释放
}

func (s *Server) AddRouter(msgID uint32, router tcpiface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router successfully!!!")
}

func (s *Server) Server() {
	//启动Server的服务功能
	s.Start()
	//TODO	做一些服务启动之后的额外的业务

	//阻塞状态
	select {}
}

/* 初始化server模块的方法

 */
func NewServer(name string) tcpiface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  utils.GlobalObject.IPVersion,
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.Port,
		MsgHandler: NewMsgHandler(),
	}
	return s
}
