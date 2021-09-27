package httpaddon

import (
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/herb-go/herbplugin"
)

var id = int64(0)

func DefaultIDGererator() string {
	v := atomic.AddInt64(&id, 1)
	return strconv.FormatInt(v, 10)
}

func (a *Addon) Create(method string, url string) *Request {
	req := &Request{
		ID:       a.IDGenerator(),
		Preset:   NewPreset(),
		Response: nil,
		Status:   StatusReady,
		addon:    a,
	}
	req.Preset.URL = url
	req.Preset.Method = method
	return req
}

type Addon struct {
	IDGenerator func() string
	Client      *http.Client
	Permission  string
	Plugin      herbplugin.Plugin
}

func Create(p herbplugin.Plugin) *Addon {
	return &Addon{
		IDGenerator: DefaultIDGererator,
		Permission:  Permission,
		Plugin:      p,
	}
}
