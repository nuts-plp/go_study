> 复制openssl的配置文件到项目所在目录下
> 找到[ CA_default ] 去掉copy_extension = copy的注释  
> 找到[ req ] 打开req_extensions = v3_req 的注释  
> 找到[ v3_req ] 添加subjectAltName = @alt_names  
> 添加新的标签 [ alt_names ] 并添加标签字段 DNS.1 = *.nuts.com      

## 生成私钥和证书
```shell
openssl genpkey -algorithm RSA -out test.key #生成私钥
```
```shell
openssl req -new -nodes -key test.key -out test.csr -days 365 -subj "/C=cn/OU=myorg/O=mycomp/CN=myname" -config ./openssl.cfg -extensions v3_req #根据私钥生成证书请求文件  
```
```shell
openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key -CAcreateserial -extfile ./openssl.cfg -extensions v3_req #生成SAN证书 pem
```
  

