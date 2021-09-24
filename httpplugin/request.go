package httpplugin

import (
	"bytes"
	"net/http"
	"net/url"
	"sort"
	"sync"

	"github.com/herb-go/herbplugin"
)

const Permission = "http"

type Request struct {
	locker   sync.Mutex
	ID       string
	Preset   *Preset
	Response *Response
	client   *http.Client
	options  herbplugin.Options
}

func (r *Request) GetID() string {
	r.locker.Lock()
	defer r.locker.Unlock()
	return r.ID
}
func (r *Request) GetURL() string {
	r.locker.Lock()
	defer r.locker.Unlock()
	return r.Preset.URL
}
func (r *Request) SetURL(url string) {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.Preset.URL = url
}
func (r *Request) GetMethod() string {
	r.locker.Lock()
	defer r.locker.Unlock()
	return r.Preset.Method
}
func (r *Request) SetMethod(method string) {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.Preset.Method = method
}
func (r *Request) GetBody() []byte {
	r.locker.Lock()
	defer r.locker.Unlock()
	return r.Preset.Body
}
func (r *Request) SetBody(body []byte) {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.Preset.Body = body
}
func (r *Request) ResetHeader() {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.Preset.Header = http.Header{}
}
func (r *Request) SetHeader(name string, value string) {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.Preset.Header.Set(name, value)
}
func (r *Request) AddHeader(name string, value string) {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.Preset.Header.Add(name, value)
}
func (r *Request) DelHeader(name string) {
	r.locker.Lock()
	defer r.locker.Unlock()

	r.Preset.Header.Del(name)
}
func (r *Request) GetHeader(name string) string {
	r.locker.Lock()
	defer r.locker.Unlock()
	return r.Preset.Header.Get(name)
}
func (r *Request) HeaderValues(name string) []string {
	r.locker.Lock()
	defer r.locker.Unlock()
	return r.Preset.Header.Values(name)
}
func (r *Request) HeaderFields() []string {
	r.locker.Lock()
	defer r.locker.Unlock()
	var result = []string{}
	for k := range r.Preset.Header {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}
func (r *Request) FinishedAt() int64 {
	r.locker.Lock()
	defer r.locker.Unlock()
	if r.Response == nil {
		panic(ErrRequestNotExecuted)
	}
	return r.Response.FinishedAt
}

func (r *Request) ResponseStatusCode() int {
	r.locker.Lock()
	defer r.locker.Unlock()
	if r.Response == nil {
		panic(ErrRequestNotExecuted)
	}
	return r.Response.StatusCode
}
func (r *Request) ResponseBody() []byte {
	r.locker.Lock()
	defer r.locker.Unlock()
	if r.Response == nil {
		panic(ErrRequestNotExecuted)
	}
	return r.Response.Body
}
func (r *Request) ResponseHeader(name string) string {
	r.locker.Lock()
	defer r.locker.Unlock()
	if r.Response == nil {
		panic(ErrRequestNotExecuted)
	}
	return r.Response.Header.Get(name)
}
func (r *Request) ResponseHeaderValues(name string) []string {
	r.locker.Lock()
	defer r.locker.Unlock()
	if r.Response == nil {
		panic(ErrRequestNotExecuted)
	}
	return r.Response.Header.Values(name)
}
func (r *Request) ResponseHeaderFields() []string {
	r.locker.Lock()
	defer r.locker.Unlock()
	if r.Response == nil {
		panic(ErrRequestNotExecuted)
	}
	var result = []string{}
	for k := range r.Response.Header {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}
func (r *Request) IsExecuted() bool {
	r.locker.Lock()
	defer r.locker.Unlock()
	return r.Response != nil
}
func (r *Request) MustAsyncExecute(callback func(error)) {
	go func() {
		defer func() {
			var err error
			var ok bool
			r := recover()
			if r != nil {
				err, ok = r.(error)
				if !ok {
					err = nil
				}
			}
			go callback(err)
		}()
		r.MustExecute()
	}()
}
func (r *Request) MustExecute() {
	opt := r.options
	if !opt.MustAuthorizePermission(Permission) {
		panic(herbplugin.NewUnauthorizePermissionError(Permission))
	}
	u, err := url.Parse(r.Preset.URL)
	if err != nil {
		panic(err)
	}
	if !opt.MustAuthorizeDomain(u.Host) {
		panic(herbplugin.NewUnauthorizeDomainError(u.Host))
	}
	h := r.Preset.Header.Get("Host")
	if h != "" {
		if !opt.MustAuthorizeDomain(h) {
			panic(herbplugin.NewUnauthorizeDomainError(h))
		}
	}
	r.locker.Lock()
	req, err := http.NewRequest(r.Preset.Method, r.Preset.URL, bytes.NewBuffer(r.Preset.Body))
	r.locker.Unlock()
	if err != nil {
		panic(err)
	}
	resp, err := r.client.Do(req)
	if err != nil {
		panic(err)
	}
	res, err := ConvertResponse(resp)
	if err != nil {
		panic(err)
	}
	r.locker.Lock()
	r.Response = res
	r.locker.Unlock()
}
func (f *Factory) Create(opt herbplugin.Options, method string, url string) *Request {
	req := &Request{
		ID:       f.IDGenerator(),
		Preset:   NewPreset(),
		Response: nil,
		client:   f.Client,
	}
	req.Preset.URL = url
	req.Preset.Method = method
	return req
}

type Factory struct {
	IDGenerator func() string
	Client      *http.Client
}
