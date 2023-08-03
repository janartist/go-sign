package main

import (
	"crypto/sha256"
	"fmt"
	"go_sign/http"
	"go_sign/sign"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	// 将 Go 函数封装为 JavaScript 函数
	js.Global().Set("signEncode", js.FuncOf(signEncode))
	js.Global().Set("signDecode", js.FuncOf(signDecode))

	<-c
}

func signEncode(this js.Value, p []js.Value) interface{} {
	// 获取传递的对象参数
	obj := p[0]
	fmt.Println("Received Object:", obj)

	httpJs := &http.JS{
		This:   this,
		Values: p,
		GetTFunc: func(value js.Value) uint {
			return uint(value.Int())
		},
		GetNonceFunc: func(value js.Value) string {
			return value.String()
		},
	}
	manager := sign.NewManager(sign.NewHMACSigner(sha256.New, []byte(p[3].String())), httpJs)
	res, err := manager.Sign()

	jsObject := js.Global().Get("Object").New()
	jsObject.Set("signature", string(res))
	jsObject.Set("err", err.Error())
	return jsObject
}

func signDecode(this js.Value, p []js.Value) interface{} {
	// 获取传递的对象参数
	obj := p[0]
	fmt.Println("Received Object:", obj)

	httpJs := &http.JS{
		This:   this,
		Values: p,
		GetTFunc: func(value js.Value) uint {
			return uint(value.Int())
		},
		GetNonceFunc: func(value js.Value) string {
			return value.String()
		},
	}
	manager := sign.NewManager(sign.NewHMACSigner(sha256.New, []byte(p[3].String())), httpJs)
	res, err := manager.Verify([]byte(p[4].String()))

	jsObject := js.Global().Get("Object").New()
	jsObject.Set("ok", res)
	jsObject.Set("err", err.Error())
	return jsObject
}
