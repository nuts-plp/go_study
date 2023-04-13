package tcpiface

type IServer interface {
	//启动服务器
	Start()
	//终止服务器
	Stop()
	//运行服务器
	Server()
}
