package httpaddon

import (
	"bytes"
	"net/http"
	"net/url"
	"sort"
	"sync"

	"github.com/herb-go/herbplugin"
)

const Permission = "http"

const StatusReady = int(0)
const StatusExecuting = int(1)
const StatusSuccess = int(2)
const StatusFail = int(3)

type Request struct {
	locker   sync.Mutex
	ID       string
	Preset   *Preset
	Status   int
	Response *Response
	addon    *Addon
}

func (r *Request) SetProxy(p string) {
	r.locker.Lock()
	defer r.locker.Unlock()
	r.Preset.Proxy = p
}
func (r *Request) GetProxy() string {
	r.locker.Lock()
	defer r.locker.Unlock()
	return r.Preset.Proxy
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
func (r *Request) ExecuteStauts() int {
	r.locker.Lock()
	defer r.locker.Unlock()
	return r.Status
}
func (r *Request) AsyncExecute(callback func(error)) {
	go func() {
		defer func() {
			var err error
			r := recover()
			if r != nil {
				err, _ = r.(error)
			}
			go callback(err)
		}()
		r.MustExecute()
	}()
}
func (r *Request) MustExecute() {
	r.locker.Lock()
	if r.Status != StatusReady {
		r.locker.Unlock()
		panic(ErrRequestExecuted)
	}
	r.locker.Unlock()
	opt := r.addon.Plugin.PluginOptions()
	if !opt.MustAuthorizePermission(r.addon.Permission) {
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
	r.locker.Lock()
	req.Header = r.Preset.Header
	r.locker.Unlock()
	c := r.addon.Client
	if c == nil {
		c = &http.Client{}
	}
	r.locker.Lock()
	if r.Preset.Proxy != "" {
		proxyurl, err := url.Parse(r.Preset.Proxy)
		if err != nil {
			r.locker.Unlock()
			panic(err)
		}
		c.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyurl),
		}
	} else {
		c.Transport = nil
	}
	r.Status = StatusExecuting
	r.locker.Unlock()
	resp, err := c.Do(req)
	if err != nil {
		r.locker.Lock()
		r.Status = StatusFail
		r.locker.Unlock()
		panic(err)
	}
	res, err := ConvertResponse(resp)
	if err != nil {
		r.locker.Lock()
		r.Status = StatusFail
		r.locker.Unlock()
		panic(err)
	}
	r.locker.Lock()
	r.Status = StatusSuccess
	r.Response = res
	r.locker.Unlock()
}
