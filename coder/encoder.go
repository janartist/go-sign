package coder

import "net/url"

type Encoder interface {
	GetT() uint
	GetNonce() string
	ToUrlValues() (url.Values, error)
}
