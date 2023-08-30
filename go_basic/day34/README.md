 ### proto
##### 编译一般proto文件
指令
`protoc --go_out=./ *.proto`
./ 表示在当前目录下   *.proto表示所有的以proto为后缀的文件


##### 编译rpc服务
**注意:**生成的是接口 <br>
指令
`protoc --go-grpc_out=./ *.proto`