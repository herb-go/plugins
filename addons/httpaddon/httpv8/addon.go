package httpv8

import (
	"net/url"

	v8plugin "github.com/herb-go/herbplugin/v8plugin"
	v8 "rogchap.com/v8go"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/plugins/addons/httpaddon"
)

type Builder func(r *v8.Context, req *Request) *v8.Value

var DefaultBuilder = func(r *v8.Context, req *Request) *v8.Value {
	obj := v8.NewObjectTemplate(r.Isolate())
	v8plugin.MustSetObjectMethod(r, obj, "GetID", req.GetID)
	v8plugin.MustSetObjectMethod(r, obj, "GetURL", req.GetURL)
	v8plugin.MustSetObjectMethod(r, obj, "SetURL", req.SetURL)
	v8plugin.MustSetObjectMethod(r, obj, "GetProxy", req.GetProxy)
	v8plugin.MustSetObjectMethod(r, obj, "SetProxy", req.SetProxy)
	v8plugin.MustSetObjectMethod(r, obj, "GetMethod", req.GetMethod)
	v8plugin.MustSetObjectMethod(r, obj, "SetMethod", req.SetMethod)
	v8plugin.MustSetObjectMethod(r, obj, "GetBody", req.GetBody)
	v8plugin.MustSetObjectMethod(r, obj, "SetBody", req.SetBody)
	v8plugin.MustSetObjectMethod(r, obj, "FinishedAt", req.FinishedAt)
	v8plugin.MustSetObjectMethod(r, obj, "ExecuteStatus", req.ExecuteStatus)
	v8plugin.MustSetObjectMethod(r, obj, "ResetHeader", req.ResetHeader)
	v8plugin.MustSetObjectMethod(r, obj, "SetHeader", req.SetHeader)
	v8plugin.MustSetObjectMethod(r, obj, "AddHeader", req.AddHeader)
	v8plugin.MustSetObjectMethod(r, obj, "DelHeader", req.DelHeader)
	v8plugin.MustSetObjectMethod(r, obj, "GetHeader", req.GetHeader)
	v8plugin.MustSetObjectMethod(r, obj, "HeaderValues", req.HeaderValues)
	v8plugin.MustSetObjectMethod(r, obj, "HeaderFields", req.HeaderFields)
	v8plugin.MustSetObjectMethod(r, obj, "ResponseStatusCode", req.ResponseStatusCode)
	v8plugin.MustSetObjectMethod(r, obj, "ResponseBody", req.ResponseBody)
	v8plugin.MustSetObjectMethod(r, obj, "ResponseHeader", req.ResponseHeader)
	v8plugin.MustSetObjectMethod(r, obj, "ResponseHeaderValues", req.ResponseHeaderValues)
	v8plugin.MustSetObjectMethod(r, obj, "ResponseHeaderFields", req.ResponseHeaderFields)
	v8plugin.MustSetObjectMethod(r, obj, "Execute", req.Execute)
	v, err := obj.NewInstance(r)
	if err != nil {
		panic(err)
	}
	return v.Value
}

type Request struct {
	Request *httpaddon.Request
}

func (req *Request) GetID(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), req.Request.GetID())
}

func (req *Request) GetURL(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), req.Request.GetURL())
}
func (req *Request) SetURL(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	req.Request.SetURL(v8plugin.MustGetArg(call, 0).String())
	return v8.Null(call.Context().Isolate())
}
func (req *Request) GetProxy(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), req.Request.GetProxy())
}
func (req *Request) SetProxy(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	req.Request.SetProxy(v8plugin.MustGetArg(call, 0).String())
	return v8.Null(call.Context().Isolate())
}

func (req *Request) GetMethod(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), req.Request.GetMethod())
}
func (req *Request) SetMethod(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	req.Request.SetMethod(v8plugin.MustGetArg(call, 0).String())
	return v8.Null(call.Context().Isolate())
}
func (req *Request) GetBody(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), string(req.Request.GetBody()))
}
func (req *Request) SetBody(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	req.Request.SetBody([]byte(v8plugin.MustGetArg(call, 0).String()))
	return nil
}

func (req *Request) FinishedAt(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), req.Request.FinishedAt())

}
func (req *Request) ExecuteStatus(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), int32(req.Request.ExecuteStauts()))
}
func (req *Request) ResetHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	req.Request.ResetHeader()
	return nil
}
func (req *Request) SetHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	req.Request.SetHeader(v8plugin.MustGetArg(call, 0).String(), v8plugin.MustGetArg(call, 1).String())
	return nil
}
func (req *Request) AddHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	req.Request.AddHeader(v8plugin.MustGetArg(call, 0).String(), v8plugin.MustGetArg(call, 1).String())
	return nil
}
func (req *Request) DelHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	req.Request.DelHeader(v8plugin.MustGetArg(call, 0).String())
	return nil

}
func (req *Request) GetHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), req.Request.GetHeader(v8plugin.MustGetArg(call, 0).String()))

}
func (req *Request) HeaderValues(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	result := req.Request.HeaderValues(v8plugin.MustGetArg(call, 0).String())
	var output = make([]v8.Valuer, len(result))
	for i, v := range result {
		output[i] = v8plugin.MustNewValue(call.Context(), v)
	}
	return v8plugin.MustNewArray(call.Context(), output)
}
func (req *Request) HeaderFields(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	result := req.Request.HeaderFields()
	var output = make([]v8.Valuer, len(result))
	for i, v := range result {
		output[i] = v8plugin.MustNewValue(call.Context(), v)
	}
	return v8plugin.MustNewArray(call.Context(), output)

}

func (req *Request) ResponseStatusCode(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), int32(req.Request.ResponseStatusCode()))
}
func (req *Request) ResponseBody(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), string(req.Request.ResponseBody()))
}
func (req *Request) ResponseHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	return v8plugin.MustNewValue(call.Context(), req.Request.ResponseHeader(v8plugin.MustGetArg(call, 0).String()))

}
func (req *Request) ResponseHeaderValues(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	result := req.Request.ResponseHeaderValues(v8plugin.MustGetArg(call, 0).String())
	var output = make([]v8.Valuer, len(result))
	for i, v := range result {
		output[i] = v8plugin.MustNewValue(call.Context(), v)
	}
	return v8plugin.MustNewArray(call.Context(), output)
}
func (req *Request) ResponseHeaderFields(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	result := req.Request.ResponseHeaderFields()
	var output = make([]v8.Valuer, len(result))
	for i, v := range result {
		output[i] = v8plugin.MustNewValue(call.Context(), v)
	}
	return v8plugin.MustNewArray(call.Context(), output)

}
func (req *Request) Execute(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	req.Request.MustExecute()
	return nil
}

type Addon struct {
	Addon   *httpaddon.Addon
	Builder Builder
}

func (a *Addon) ParseURL(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	rawurl := call.Args()[0].String()
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil
	}
	result := v8.NewObjectTemplate(call.Context().Isolate())
	result.Set("Host", u.Host)
	result.Set("Hostname", u.Host)
	result.Set("Scheme", u.Scheme)
	result.Set("Path", u.Path)
	result.Set("Query", u.RawQuery)
	result.Set("User", u.User.Username())
	p, _ := u.User.Password()
	result.Set("Password", p)
	result.Set("Port", u.Port())
	result.Set("Fragment", u.Fragment)
	obj, err := result.NewInstance(call.Context())
	if err != nil {
		panic(err)
	}
	return obj.Value
}
func (a *Addon) NewRequest(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	method := v8plugin.MustGetArg(call, 0).String()
	url := v8plugin.MustGetArg(call, 1).String()
	req := a.Addon.Create(method, url)
	return a.Builder(call.Context(), &Request{req})
}

func (a *Addon) Convert(r *v8.Context) *v8.Value {
	obj := v8.NewObjectTemplate(r.Isolate())
	v8plugin.MustSetObjectMethod(r, obj, "New", a.NewRequest)
	v8plugin.MustSetObjectMethod(r, obj, "ParseURL", a.ParseURL)
	return v8plugin.MustObjectTemplateToValue(r, obj)
}
func Create(p herbplugin.Plugin) *Addon {
	return &Addon{
		Addon:   httpaddon.Create(p),
		Builder: DefaultBuilder,
	}
}
