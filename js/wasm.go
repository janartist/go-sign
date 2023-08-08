package js

import (
	"crypto/sha256"
	"fmt"
	"syscall/js"

	"github.com/janartist/go-sign/sign"
)

type wasm struct {
}

func NewWASM() *wasm {
	return &wasm{}
}
func (wasm) SignEncode(this js.Value, p []js.Value) interface{} {
	defer func() {
		if r := recover(); r != nil {
			// 处理捕获的异常
			console := js.Global().Get("console")
			console.Call("error", "Caught an error:", r)
		}
	}()
	// 获取传递的对象参数
	fmt.Println("Received signEncode this:", this)
	fmt.Println("Received signEncode p:", p)

	jsObject := js.Global().Get("Object").New()

	if len(p) != 2 {
		jsObject.Set("err", "p len is not 2")
		return jsObject
	}
	if p[1].Type() != js.TypeObject {
		jsObject.Set("err", "p[1] is not obj")
		return jsObject
	}

	httpJs := &JS{
		This:   this,
		Values: p,
		GetTFunc: func(values []js.Value) js.Value {
			return values[1].Get("t")
		},
		GetNonceFunc: func(values []js.Value) js.Value {
			return values[1].Get("nonce")
		},
		GetValuesFunc: func(values []js.Value) js.Value {
			return values[0]
		},
	}
	manager := sign.NewManager(sign.NewHMACSigner(sha256.New, []byte(p[1].Get("secret").String())), httpJs)
	res, str, err := manager.Sign()
	jsObject.Set("signature", string(res))
	jsObject.Set("str", str)
	err2 := ""
	if err != nil {
		err2 = err.Error()
	}
	jsObject.Set("err", err2)
	return jsObject
}

func (wasm) SignDecode(this js.Value, p []js.Value) interface{} {
	defer func() {
		if r := recover(); r != nil {
			// 处理捕获的异常
			console := js.Global().Get("console")
			console.Call("error", "Caught an error:", r)
		}
	}()
	// 获取传递的对象参数
	fmt.Println("Received signDecode this:", this)
	fmt.Println("Received signDecode p:", p)

	jsObject := js.Global().Get("Object").New()
	jsObject.Set("ok", false)

	if len(p) != 2 {
		jsObject.Set("err", "p len is not 2")
		return jsObject
	}
	if p[1].Type() != js.TypeObject {
		jsObject.Set("err", "p[1] is not obj")
		return jsObject
	}

	httpJs := &JS{
		This:   this,
		Values: p,
		GetTFunc: func(values []js.Value) js.Value {
			return values[1].Get("t")
		},
		GetNonceFunc: func(values []js.Value) js.Value {
			return values[1].Get("nonce")
		},
		GetValuesFunc: func(values []js.Value) js.Value {
			return values[0]
		},
	}
	manager := sign.NewManager(sign.NewHMACSigner(sha256.New, []byte(p[1].Get("secret").String())), httpJs)
	ok, str, err := manager.Verify([]byte(p[0].Get("signature").String()))

	jsObject.Set("ok", ok)
	jsObject.Set("str", str)
	err2 := ""
	if err != nil {
		err2 = err.Error()
	}
	jsObject.Set("err", err2)
	return jsObject
}

func (wasm) TestAlert(this js.Value, p []js.Value) interface{} {
	defer func() {
		if r := recover(); r != nil {
			// 处理捕获的异常
			console := js.Global().Get("console")
			console.Call("error", "Caught an error:", r)
		}
	}()
	// 调用测试函数
	result := "ok！"
	// 返回结果
	js.Global().Call("alert", fmt.Sprintf("Result: %s", result))

	return result
}
