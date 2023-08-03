package sign

import (
	"go_sign/coder"
	"strconv"
)

const (
	paramT     = "t"
	paramNonce = "nonce"
)

func NewManager(signer Signer, coder coder.Coder) *Manager {
	return &Manager{
		signer: signer,
		coder:  coder,
	}
}

type Manager struct {
	signer Signer
	coder  coder.Coder
}

func (m *Manager) Sign() ([]byte, error) {
	values, err := m.coder.ToUrlValues()
	if err != nil {
		return nil, err
	}
	values.Set(paramT, strconv.Itoa(int(m.coder.GetT())))
	values.Set(paramNonce, m.coder.GetNonce())
	return m.signer.Sign([]byte(values.Encode()))
}

func (m *Manager) Verify(signature []byte) (bool, error) {
	values, err := m.coder.ToUrlValues()
	if err != nil {
		return false, err
	}
	values.Set(paramT, strconv.Itoa(int(m.coder.GetT())))
	values.Set(paramNonce, m.coder.GetNonce())
	return m.signer.Verify([]byte(values.Encode()), signature)
}
