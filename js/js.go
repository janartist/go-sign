package js

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/janartist/go-sign/sign"
	"net/url"
	"syscall/js"

	"github.com/janartist/go-sign/tool"
)

type JS struct {
	This          js.Value
	Values        []js.Value
	GetTFunc      func([]js.Value) js.Value
	GetNonceFunc  func([]js.Value) js.Value
	GetValuesFunc func([]js.Value) js.Value
}

func (j *JS) GetT() uint {
	return uint(j.GetTFunc(j.Values).Int())
}

func (j *JS) GetNonce() string {
	return j.GetNonceFunc(j.Values).String()
}

func (j *JS) ToUrlValues() (url.Values, error) {
	// 将对象的属性提取为 URL 查询参数
	val := j.GetValuesFunc(j.Values)
	urlValues := tool.FlattenData(convertJSValue(val))
	return urlValues, nil
}

func convertJSValue(value js.Value) interface{} {
	switch value.Type() {
	case js.TypeUndefined, js.TypeNull:
		return nil
	case js.TypeBoolean:
		return value.Bool()
	case js.TypeNumber:
		return value.Float()
	case js.TypeString:
		return value.String()
	case js.TypeObject:
		if isArray(value) {
			return convertJSArray(value)
		}
		return convertJSObject(value)
	default:
		return nil
	}
}

func convertJSArray(value js.Value) []interface{} {
	length := value.Length()
	result := make([]interface{}, length)
	for i := 0; i < length; i++ {
		result[i] = convertJSValue(value.Index(i))
	}
	return result
}

func convertJSObject(value js.Value) map[string]interface{} {
	result := make(map[string]interface{})
	value.Call("constructor", js.Global().Get("Object")).Call("keys", value).Call("forEach", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		key := p[0].String()
		result[key] = convertJSValue(value.Get(key))
		return nil
	}))
	return result
}

func isArray(value js.Value) bool {
	return value.InstanceOf(js.Global().Get("Array"))
}

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
	utf8Str := string(res)
	fmt.Println("signature:", utf8Str)
	jsObject.Set("signature", base64.StdEncoding.EncodeToString([]byte(utf8Str)))
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
