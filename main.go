package main

import (
	"syscall/js"

	js2 "github.com/janartist/go-sign/js"
)

func main() {
	wasm := js2.NewWASM()
	// 将 Go 函数封装为 JavaScript 函数
	js.Global().Set("signEncode", js.FuncOf(wasm.SignEncode))
	js.Global().Set("signDecode", js.FuncOf(wasm.SignDecode))
	js.Global().Set("testAlert", js.FuncOf(wasm.TestAlert))

	<-make(chan struct{})
}
