package coder

import "net/url"

type Coder interface {
	GetT() uint
	GetNonce() string
	// Encode() ([]byte, error)
	ToUrlValues() (url.Values, error)
}
