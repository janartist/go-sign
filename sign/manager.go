package sign

import (
	"strconv"

	"github.com/janartist/go-sign/coder"
)

const (
	paramT     = "t"
	paramNonce = "nonce"
)

func NewManager(signer Signer, coder coder.Encoder) *Manager {
	return &Manager{
		signer: signer,
		coder:  coder,
	}
}

type Manager struct {
	signer Signer
	coder  coder.Encoder
}

// Sign  统一签名方法
// signature为数据 str为签名前有序字符串
func (m *Manager) Sign() (signature []byte, str string, err error) {
	values, err := m.coder.ToUrlValues()
	if err != nil {
		return
	}
	values.Add(paramT, strconv.Itoa(int(m.coder.GetT())))
	values.Add(paramNonce, m.coder.GetNonce())
	str = values.Encode()
	signature, err = m.signer.Sign([]byte(str))
	return
}

// Verify  统一签名方法 signature为数据
// str为签名前有序字符串
func (m *Manager) Verify(signature []byte) (ok bool, str string, err error) {
	values, err := m.coder.ToUrlValues()
	if err != nil {
		return
	}
	values.Add(paramT, strconv.Itoa(int(m.coder.GetT())))
	values.Add(paramNonce, m.coder.GetNonce())
	str = values.Encode()
	ok, err = m.signer.Verify([]byte(str), signature)
	return
}
