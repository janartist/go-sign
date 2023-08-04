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
* 签名方法
```js
// 示例数据对象，必须为对象，数据随意，支持无限级
const data = {
    value: 42,
    val:"wo de我",
};

// 固定示例配置对象
const config = {
    t: 12345, //时间戳
    nonce: "wefw23swdwef", //随机数
    secret: 'mysecret', //密钥，视算法定，现为hmac-sha256
};
var res = signEncode.call(this, data, config);
console.log(res)
{
    "signature": "MTZlNGJkNmU3NThkM2FlMjRjM2U0ZDVkOTU5MTYxMjVkMWRlZDllZDg1OGNjZDBiNjliMTQwODFkYjhkNmE0NA==",
    "str": "nonce=wefw23swdwef&t=12345&val=wo+de%E6%88%91&value=42",
    "err": ""
}
```

### grpc示例
开发中