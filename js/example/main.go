package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("./index.html"); os.IsNotExist(err) {
			http.Error(w, "index.html not found", http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, "./index.html")
	})
	http.HandleFunc("/wasm.js", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("./wasm.js"); os.IsNotExist(err) {
			http.Error(w, "wasm_exec.js not found", http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, "./wasm.js")
	})
	http.HandleFunc("/wasm_exec.js", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("./wasm_exec.js"); os.IsNotExist(err) {
			http.Error(w, "wasm_exec.js not found", http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, "./wasm_exec.js")
	})
	http.HandleFunc("/main.wasm", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("./main.wasm"); os.IsNotExist(err) {
			http.Error(w, "main.wasm not found", http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, "./main.wasm")
	})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
