package coder

import "net/url"

type Coder interface {
	GetT() uint
	GetNonce() string
	ToUrlValues() (url.Values, error)
}
