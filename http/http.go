package http

import (
	"github.com/janartist/go-sign/tool"

	"encoding/json"
	"net/http"
	"net/url"
)

const (
	contentTypeJson = "application/json"
)

type Request struct {
	*http.Request
	GetTFunc     func(*http.Request) uint
	GetNonceFunc func(*http.Request) string
}

func (r *Request) GetT() uint {
	return r.GetTFunc(r.Request)
}

func (r *Request) GetNonce() string {
	return r.GetNonceFunc(r.Request)
}

func (r *Request) ToUrlValues() (url.Values, error) {
	urlValues := make(url.Values)
	if r.Header.Get("Content-Type") == contentTypeJson {
		var data interface{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return nil, err
		}
		urlValues = tool.FlattenData(data)
	} else {
		if len(r.Form) == 0 {
			err := r.ParseForm()
			if err != nil {
				return nil, err
			}
			urlValues = r.Form
		}
	}
	return urlValues, nil
}
