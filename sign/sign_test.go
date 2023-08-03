package sign_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"go_sign/sign"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []byte("我是内容12138")

func TestHMACSigner(t *testing.T) {
	secretKey := []byte("se1")
	secretKey2 := []byte("se1x")
	t.Run("sha256", func(t *testing.T) {
		MACSigner := sign.NewHMACSigner(sha256.New, secretKey)
		MACSigner2 := sign.NewHMACSigner(sha256.New, secretKey2)
		signature, err := MACSigner.Sign(data)
		assert.Nil(t, err)
		ok, err := MACSigner.Verify(data, signature)
		assert.Nil(t, err)
		assert.Equal(t, ok, true)
		ok, err = MACSigner2.Verify(data, signature)
		assert.Nil(t, err)
		assert.Equal(t, ok, false)
	})
	t.Run("md5", func(t *testing.T) {
		MACSigner := sign.NewHMACSigner(md5.New, secretKey)
		MACSigner2 := sign.NewHMACSigner(sha256.New, secretKey2)
		signature, err := MACSigner.Sign(data)
		assert.Nil(t, err)
		ok, err := MACSigner.Verify(data, signature)
		assert.Nil(t, err)
		assert.Equal(t, ok, true)
		ok, err = MACSigner2.Verify(data, signature)
		assert.Nil(t, err)
		assert.Equal(t, ok, false)
	})
}

func aTestRSASigner(t *testing.T) {
	// 读取 RSA 私钥
	rsaKeyData, err := ioutil.ReadFile("./private.key")
	assert.Nil(t, err)

	block, _ := pem.Decode(rsaKeyData)
	assert.NotNil(t, block)
	assert.Equal(t, block.Type, "PRIVATE KEY")

	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	assert.NotNil(t, rsaPrivateKey)
	assert.Nil(t, err)
	// 读取 RSA 公钥
	rsaKeyData, err = ioutil.ReadFile("./public.key")
	assert.Nil(t, err)
	block, _ = pem.Decode(rsaKeyData)
	assert.NotNil(t, block)
	assert.Equal(t, block.Type, "PUBLIC KEY")
	rsaPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	assert.Nil(t, err)
	assert.NotNil(t, rsaPublicKey)

	RSASignerPrivateKey := sign.NewRSASigner(rsaPrivateKey)
	RSASignerPublicKey := sign.NewRSASigner(&rsa.PrivateKey{PublicKey: *rsaPublicKey.(*rsa.PublicKey)})
	signature, err := RSASignerPrivateKey.Sign(data)
	assert.Nil(t, err)
	ok, err := RSASignerPrivateKey.Verify(data, signature)
	assert.Nil(t, err)
	assert.Equal(t, ok, true)
	ok, err = RSASignerPublicKey.Verify(data, signature)
	assert.Nil(t, err)
	assert.Equal(t, ok, true)
	fmt.Print(1122)
}

func TestECDSASigner(t *testing.T) {
	ecdsaPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if !assert.Nil(t, err) {
		return
	}

	ECDSASigner := sign.NewECDSASigner(ecdsaPrivateKey)
	signature, err := ECDSASigner.Sign(data)
	assert.Nil(t, err)
	ok, err := ECDSASigner.Verify(data, signature)
	assert.Nil(t, err)
	assert.Equal(t, ok, true)
}
