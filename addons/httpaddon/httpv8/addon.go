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
	obj.Set("GetID", v8.NewFunctionTemplate(r.Isolate(), req.GetID))
	obj.Set("GetURL", v8.NewFunctionTemplate(r.Isolate(), req.GetURL))
	obj.Set("SetURL", v8.NewFunctionTemplate(r.Isolate(), req.SetURL))
	obj.Set("GetProxy", v8.NewFunctionTemplate(r.Isolate(), req.GetProxy))
	obj.Set("SetProxy", v8.NewFunctionTemplate(r.Isolate(), req.SetProxy))
	obj.Set("GetMethod", v8.NewFunctionTemplate(r.Isolate(), req.GetMethod))
	obj.Set("SetMethod", v8.NewFunctionTemplate(r.Isolate(), req.SetMethod))
	obj.Set("GetBody", v8.NewFunctionTemplate(r.Isolate(), req.GetBody))
	obj.Set("SetBody", v8.NewFunctionTemplate(r.Isolate(), req.SetBody))
	obj.Set("FinishedAt", v8.NewFunctionTemplate(r.Isolate(), req.FinishedAt))
	obj.Set("ExecuteStatus", v8.NewFunctionTemplate(r.Isolate(), req.ExecuteStatus))
	obj.Set("ResetHeader", v8.NewFunctionTemplate(r.Isolate(), req.ResetHeader))
	obj.Set("SetHeader", v8.NewFunctionTemplate(r.Isolate(), req.SetHeader))
	obj.Set("AddHeader", v8.NewFunctionTemplate(r.Isolate(), req.AddHeader))
	obj.Set("DelHeader", v8.NewFunctionTemplate(r.Isolate(), req.DelHeader))
	obj.Set("GetHeader", v8.NewFunctionTemplate(r.Isolate(), req.GetHeader))
	obj.Set("HeaderValues", v8.NewFunctionTemplate(r.Isolate(), req.HeaderValues))
	obj.Set("HeaderFields", v8.NewFunctionTemplate(r.Isolate(), req.HeaderFields))
	obj.Set("ResponseStatusCode", v8.NewFunctionTemplate(r.Isolate(), req.ResponseStatusCode))
	obj.Set("ResponseBody", v8.NewFunctionTemplate(r.Isolate(), req.ResponseBody))
	obj.Set("ResponseHeader", v8.NewFunctionTemplate(r.Isolate(), req.ResponseHeader))
	obj.Set("ResponseHeaderValues", v8.NewFunctionTemplate(r.Isolate(), req.ResponseHeaderValues))
	obj.Set("ResponseHeaderFields", v8.NewFunctionTemplate(r.Isolate(), req.ResponseHeaderFields))
	obj.Set("Execute", v8.NewFunctionTemplate(r.Isolate(), req.Execute))
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
	return v8plugin.MustNewValue(call.Context(), req.Request.GetID())
}

func (req *Request) GetURL(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), req.Request.GetURL())
}
func (req *Request) SetURL(call *v8.FunctionCallbackInfo) *v8.Value {
	req.Request.SetURL(v8plugin.MustGetArg(call, 0).String())
	return v8.Null(call.Context().Isolate())
}
func (req *Request) GetProxy(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), req.Request.GetProxy())
}
func (req *Request) SetProxy(call *v8.FunctionCallbackInfo) *v8.Value {
	req.Request.SetProxy(v8plugin.MustGetArg(call, 0).String())
	return v8.Null(call.Context().Isolate())
}

func (req *Request) GetMethod(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), req.Request.GetMethod())
}
func (req *Request) SetMethod(call *v8.FunctionCallbackInfo) *v8.Value {
	req.Request.SetMethod(v8plugin.MustGetArg(call, 0).String())
	return v8.Null(call.Context().Isolate())
}
func (req *Request) GetBody(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), string(req.Request.GetBody()))
}
func (req *Request) SetBody(call *v8.FunctionCallbackInfo) *v8.Value {
	req.Request.SetBody([]byte(v8plugin.MustGetArg(call, 0).String()))
	return nil
}

func (req *Request) FinishedAt(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), req.Request.FinishedAt())

}
func (req *Request) ExecuteStatus(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), int32(req.Request.ExecuteStauts()))
}
func (req *Request) ResetHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	req.Request.ResetHeader()
	return nil
}
func (req *Request) SetHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	req.Request.SetHeader(v8plugin.MustGetArg(call, 0).String(), v8plugin.MustGetArg(call, 1).String())
	return nil
}
func (req *Request) AddHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	req.Request.AddHeader(v8plugin.MustGetArg(call, 0).String(), v8plugin.MustGetArg(call, 1).String())
	return nil
}
func (req *Request) DelHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	req.Request.DelHeader(v8plugin.MustGetArg(call, 0).String())
	return nil

}
func (req *Request) GetHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), req.Request.GetHeader(v8plugin.MustGetArg(call, 0).String()))

}
func (req *Request) HeaderValues(call *v8.FunctionCallbackInfo) *v8.Value {
	result := req.Request.HeaderValues(v8plugin.MustGetArg(call, 0).String())
	var output = make([]v8.Valuer, len(result))
	for i, v := range result {
		output[i] = v8plugin.MustNewValue(call.Context(), v)
	}
	return v8plugin.MustNewArray(call.Context(), output)
}
func (req *Request) HeaderFields(call *v8.FunctionCallbackInfo) *v8.Value {
	result := req.Request.HeaderFields()
	var output = make([]v8.Valuer, len(result))
	for i, v := range result {
		output[i] = v8plugin.MustNewValue(call.Context(), v)
	}
	return v8plugin.MustNewArray(call.Context(), output)

}

func (req *Request) ResponseStatusCode(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), int32(req.Request.ResponseStatusCode()))
}
func (req *Request) ResponseBody(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), string(req.Request.ResponseBody()))
}
func (req *Request) ResponseHeader(call *v8.FunctionCallbackInfo) *v8.Value {
	return v8plugin.MustNewValue(call.Context(), req.Request.ResponseHeader(v8plugin.MustGetArg(call, 0).String()))

}
func (req *Request) ResponseHeaderValues(call *v8.FunctionCallbackInfo) *v8.Value {
	result := req.Request.ResponseHeaderValues(v8plugin.MustGetArg(call, 0).String())
	var output = make([]v8.Valuer, len(result))
	for i, v := range result {
		output[i] = v8plugin.MustNewValue(call.Context(), v)
	}
	return v8plugin.MustNewArray(call.Context(), output)
}
func (req *Request) ResponseHeaderFields(call *v8.FunctionCallbackInfo) *v8.Value {
	result := req.Request.ResponseHeaderFields()
	var output = make([]v8.Valuer, len(result))
	for i, v := range result {
		output[i] = v8plugin.MustNewValue(call.Context(), v)
	}
	return v8plugin.MustNewArray(call.Context(), output)

}
func (req *Request) Execute(call *v8.FunctionCallbackInfo) *v8.Value {
	req.Request.MustExecute()
	return nil
}

type Addon struct {
	Addon   *httpaddon.Addon
	Builder Builder
}

func (a *Addon) ParseURL(call *v8.FunctionCallbackInfo) *v8.Value {
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
	method := v8plugin.MustGetArg(call, 0).String()
	url := v8plugin.MustGetArg(call, 1).String()
	req := a.Addon.Create(method, url)
	return a.Builder(call.Context(), &Request{req})
}

func (a *Addon) Convert(r *v8.Context) *v8.Value {
	obj := v8.NewObjectTemplate(r.Isolate())
	obj.Set("New", v8.NewFunctionTemplate(r.Isolate(), a.NewRequest))
	obj.Set("ParseURL", v8.NewFunctionTemplate(r.Isolate(), a.ParseURL))
	return v8plugin.MustObjectTemplateToValue(obj, r)
}
func Create(p herbplugin.Plugin) *Addon {
	return &Addon{
		Addon:   httpaddon.Create(p),
		Builder: DefaultBuilder,
	}
}
