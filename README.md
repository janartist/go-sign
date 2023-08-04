## 签名解决方法

### 好处
* 数据完整性验证： 数字签名可以用来验证数据的完整性，确保数据在传输或存储过程中没有被篡改。
* 身份认证： 数字签名可以用来验证数据的发送者身份，确保接收者可以确定数据来自可信的发送者。

### 算法



支持签名：
- [HMAC](https://github.com/golang/go/blob/master/src/crypto/hmac/hmac.go)
- [RSA](https://github.com/golang/go/blob/master/src/crypto/rsa/pkcs1v15.go) 
- [ECDSA](https://github.com/golang/go/blob/master/src/crypto/ecdsa/ecdsa.go)

支持 ```http``` ```grpc``` ```js``` 端调用


### http示例

[http客户端，服务端](https://github.com/janartist/go-sign/blob/main/sign/manager_test.go)

### JS通过WebAssembly交互示例
[js调用go签名](https://github.com/janartist/go-sign/tree/main/js/example)
* 关于wasm文件编译(仓库wasm文件已编译好)
```shell
GOARCH=wasm GOOS=js go build -o js/example/main.wasm
```

### grpc示例
开发中