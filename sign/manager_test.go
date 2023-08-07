package sign_test

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	http2 "github.com/janartist/go-sign/http"
	"github.com/janartist/go-sign/sign"
	"github.com/stretchr/testify/assert"
)

var signatureStr string

func TestManager_SignVerify(t *testing.T) {
	t.Run("sign", func(t *testing.T) {
		manager_Sign(t)

	})
	t.Run("verify", func(t *testing.T) {
		manager_Verify(t)

	})
}

func manager_Sign(t *testing.T) {
	secretKey := []byte("se1")
	// 构造 HTTP 请求的内容
	url := "/posts"
	method := "POST"
	payload := []byte(`{"title":"foo我","body":"bar","userId":1,"ss":[1,23]}`)

	// 创建请求对象
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Sign-T", "1234522")
	req.Header.Set("Sign-Nonce", "wfww2211")
	manager := sign.NewManager(sign.NewHMACSigner(sha256.New, secretKey), &http2.Request{
		Request: req,
		GetTFunc: func(request *http.Request) uint {
			t, _ := strconv.Atoi(request.Header.Get("Sign-T"))
			return uint(t)
		},
		GetNonceFunc: func(request *http.Request) string {
			return request.Header.Get("Sign-Nonce")
		},
	})
	res, str, err := manager.Sign()
	fmt.Println(string(res), str)
	assert.Nil(t, err)
	signatureStr = string(res)
}

func manager_Verify(t *testing.T) {
	secretKey := []byte("se1")
	// 构造 HTTP 请求的内容
	url := "/posts"
	method := "POST"
	payload := []byte(`{"title":"foo我","body":"bar","userId":1,"ss":[1,23]}`)

	// 创建请求对象
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Sign-T", "1234522")
	req.Header.Set("Sign-Nonce", "wfww2211")
	manager := sign.NewManager(sign.NewHMACSigner(sha256.New, secretKey), &http2.Request{
		Request: req,
		GetTFunc: func(request *http.Request) uint {
			t, _ := strconv.Atoi(request.Header.Get("Sign-T"))
			return uint(t)
		},
		GetNonceFunc: func(request *http.Request) string {
			return request.Header.Get("Sign-Nonce")
		},
	})
	ok, _, err := manager.Verify([]byte(signatureStr))
	assert.Nil(t, err)
	assert.Equal(t, ok, true)
}
