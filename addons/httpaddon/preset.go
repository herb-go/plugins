package httpaddon

import "net/http"

type Preset struct {
	Body   []byte
	Method string
	Header http.Header
	URL    string
	Proxy  string
}

func NewPreset() *Preset {
	return &Preset{
		Body:   nil,
		Method: "",
		Header: http.Header{},
		URL:    "",
		Proxy:  "",
	}
}
