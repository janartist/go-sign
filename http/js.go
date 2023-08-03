package http

import (
	"go_sign/tool"
	"net/url"
	"syscall/js"
)

type JS struct {
	This         js.Value
	Values       []js.Value
	GetTFunc     func(js.Value) uint
	GetNonceFunc func(js.Value) string
}

func (j *JS) GetT() uint {
	return j.GetTFunc(j.Values[1])
}

func (j *JS) GetNonce() string {
	return j.GetNonceFunc(j.Values[2])
}

func (j *JS) ToUrlValues() (url.Values, error) {
	// 将对象的属性提取为 URL 查询参数
	val := j.Values[0]
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
