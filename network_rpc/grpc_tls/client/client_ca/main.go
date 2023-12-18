package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"strconv"

	"google.golang.org/grpc/credentials"

	grpc_tls "github.com/grpc_tls/proto"

	"google.golang.org/grpc"
)

const PORT = 9091

func main() {
	cert, err := tls.LoadX509KeyPair("./../../conf/tls_ca/client.pem", "./../../conf/tls_ca/client.key")
	if err != nil {
		log.Printf("[client] load key pairs failed! err:%s\n")
		return
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./../../conf/tls_ca/ca.crt")
	if err != nil {
		log.Printf("[client] read ca file failed! err:%s\n", err)
		return
	}
	ok := certPool.AppendCertsFromPEM(ca)
	if !ok {
		log.Printf("[client] append certs failed!  err:%s\n")
		return
	}
	newTLS := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书
		ServerName: "*.bytefree.com",
		RootCAs:    certPool,
	})
	if err != nil {
		log.Printf("[client] generate tls failed! err:%s\n", err)
		return
	}
	conn, err := grpc.Dial(":"+strconv.Itoa(PORT), grpc.WithTransportCredentials(newTLS))
	//conn, err := grpc.Dial(":"+strconv.Itoa(PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("connection create failed! %s\n", err)
		return
	}
	client := grpc_tls.NewSearchCClient(conn)
	resp, err := client.Search(context.Background(), &grpc_tls.Req{
		Request: "潘丽萍",
	})
	if err != nil {
		log.Printf("[client] rpc invoke failed! err:%s\n", err)
		return
	}
	log.Printf("[client] resp value:%s\n", resp.GetResponse())
}
