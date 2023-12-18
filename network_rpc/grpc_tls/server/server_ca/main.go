package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"

	grpc_tls "github.com/grpc_tls/proto"
)

const PORT = 9091

type Serv struct {
	grpc_tls.UnimplementedSearchCServer
}

func (s *Serv) Search(ctx context.Context, req *grpc_tls.Req) (*grpc_tls.Resp, error) {
	return &grpc_tls.Resp{
		Response: "hello!" + req.GetRequest() + ",你上当了！",
	}, nil

}
func main() {

	cert, err := tls.LoadX509KeyPair("./../../conf/tls_ca/server.pem", "./../../conf/tls_ca/server.key")
	if err != nil {
		log.Printf("[server] load ca file failed! err:%s\n", err)
		return
	}
	//创建一个新的、空的CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./../../conf/tls_ca/ca.crt")
	if err != nil {
		log.Printf("[server] read ca failed! err:%s\n", err)
		return
	}
	//尝试解析传入的pem编码的证书。如果解析成功将其添加到CertPool中以便后面的使用
	ok := certPool.AppendCertsFromPEM(ca)
	if !ok {
		log.Printf("[server] certPoll append certs failed! err:%s\n", err)
		return
	}
	//构建基于TLS的 TransportCredentials 选项
	c := credentials.NewTLS(&tls.Config{
		//设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		//要求必须校验客户端证书
		ClientAuth: tls.RequireAndVerifyClientCert,
		//设置根证书集合，校验方式选择ClientAuth 中设定的模式
		ClientCAs: certPool,
	})
	//监听
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Printf("[server] listen %d failed!", err)
		return
	}
	//server := grpc.NewServer()
	server := grpc.NewServer(grpc.Creds(c))
	grpc_tls.RegisterSearchCServer(server, &Serv{})
	_ = server.Serve(conn)

}
