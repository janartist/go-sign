package sign

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/asn1"
	"hash"
	"math/big"
)

// Signer 是一个数字签名接口
type Signer interface {
	Sign(data []byte) ([]byte, error)
	Verify(data, signature []byte) (bool, error)
}

// hMACSigner 支持 HMAC 签名
type hMACSigner struct {
	hashFunc  func() hash.Hash
	secretKey []byte
}

func NewHMACSigner(hashFunc func() hash.Hash, secretKey []byte) *hMACSigner {
	return &hMACSigner{
		hashFunc:  hashFunc,
		secretKey: secretKey,
	}
}

func (h *hMACSigner) Sign(data []byte) ([]byte, error) {
	// 创建一个 HMAC 对象
	hc := hmac.New(h.hashFunc, h.secretKey)
	// 写入消息
	hc.Write(data)

	// 计算 HMAC 值
	signature := hc.Sum(nil)
	return signature, nil
}

func (h *hMACSigner) Verify(data, signature []byte) (bool, error) {
	sign, err := h.Sign(data)
	if err != nil {
		return false, err
	}
	return hmac.Equal(sign, signature), nil
}

// rSASigner 支持 RSA 签名
type rSASigner struct {
	privateKey *rsa.PrivateKey
}

func NewRSASigner(privateKey *rsa.PrivateKey) *rSASigner {
	return &rSASigner{
		privateKey: privateKey,
	}
}

func (r *rSASigner) Sign(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	return rsa.SignPKCS1v15(rand.Reader, r.privateKey, crypto.SHA256, hash[:])
}
func (r *rSASigner) Verify(data, signature []byte) (bool, error) {
	hash := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(&r.privateKey.PublicKey, crypto.SHA256, hash[:], signature)
	return err == nil, err
}

// eCDSASigner 支持 ECDSA 签名
type eCDSASigner struct {
	privateKey *ecdsa.PrivateKey
}

func NewECDSASigner(privateKey *ecdsa.PrivateKey) *eCDSASigner {
	return &eCDSASigner{
		privateKey: privateKey,
	}
}

func (e *eCDSASigner) Sign(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	return ecdsa.SignASN1(rand.Reader, e.privateKey, hash[:])
}

func (e *eCDSASigner) Verify(data, signature []byte) (bool, error) {
	hash := sha256.Sum256(data)
	// 解码签名以获取 r 和 s
	var ecdsaSig struct {
		R, S *big.Int
	}
	_, err := asn1.Unmarshal(signature, &ecdsaSig)
	if err != nil {
		return false, err
	}

	return ecdsa.Verify(&e.privateKey.PublicKey, hash[:], ecdsaSig.R, ecdsaSig.S), nil
}
