package tcpnet

import "zinx/tcpiface"

type Request struct {
	//已经和客户端建立好的链接
	conn tcpiface.IConnection

	//客户端请求的数据
	data []byte
}

//得到当前链接
func (r *Request) GetConnection() tcpiface.IConnection {
	return r.conn
}

//获取请求数据
func (r *Request) GetData() []byte {
	return r.data
}
