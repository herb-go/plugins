package httpjsplugin

import (
	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/plugins/addons/httpaddon"
)

type Builder func(r *goja.Runtime, req *Request) *goja.Object

var DefaultBuilder = func(r *goja.Runtime, req *Request) *goja.Object {
	obj := r.NewObject()
	obj.Set("GetID", req.GetID)
	obj.Set("GetURL", req.GetURL)
	obj.Set("SetURL", req.SetURL)
	obj.Set("GetMethod", req.GetMethod)
	obj.Set("SetMethod", req.SetMethod)
	obj.Set("GetBody", req.GetBody)
	obj.Set("SetBody", req.SetBody)
	obj.Set("FinishedAt", req.FinishedAt)
	obj.Set("ExecuteStatus", req.ExecuteStatus)
	obj.Set("ResetHeader", req.ResetHeader)
	obj.Set("SetHeader", req.SetHeader)
	obj.Set("AddHeader", req.AddHeader)
	obj.Set("DelHeader", req.DelHeader)
	obj.Set("GetHeader", req.GetHeader)
	obj.Set("HeaderValues", req.HeaderValues)
	obj.Set("HeaderFields", req.HeaderFields)
	obj.Set("ResponseStatusCode", req.ResponseStatusCode)
	obj.Set("ResponseBody", req.ResponseBody)
	obj.Set("ResponseHeader", req.ResponseHeader)
	obj.Set("ResponseHeaderValues", req.ResponseHeaderValues)
	obj.Set("ResponseHeaderFields", req.ResponseHeaderFields)
	obj.Set("Execute", req.Execute)
	return obj
}

type Request struct {
	Request *httpaddon.Request
}

func (req *Request) GetID(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(req.Request.GetID())
}

func (req *Request) GetURL(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(req.Request.GetURL())
}
func (req *Request) SetURL(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	req.Request.SetURL(call.Argument(0).String())
	return nil
}
func (req *Request) GetMethod(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(req.Request.GetMethod())
}
func (req *Request) SetMethod(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	req.Request.SetMethod(call.Argument(0).String())
	return nil

}
func (req *Request) GetBody(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(string(req.Request.GetBody()))
}
func (req *Request) SetBody(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	req.Request.SetBody([]byte(call.Argument(0).String()))
	return nil

}
func (req *Request) FinishedAt(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(req.Request.FinishedAt())

}
func (req *Request) ExecuteStatus(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(req.Request.ExecuteStauts())
}
func (req *Request) ResetHeader(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	req.Request.ResetHeader()
	return nil
}
func (req *Request) SetHeader(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	req.Request.SetHeader(call.Argument(0).String(), call.Argument(1).String())
	return nil
}
func (req *Request) AddHeader(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	req.Request.AddHeader(call.Argument(0).String(), call.Argument(1).String())
	return nil
}
func (req *Request) DelHeader(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	req.Request.DelHeader(call.Argument(0).String())
	return nil

}
func (req *Request) GetHeader(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(req.Request.GetHeader(call.Argument(0).String()))

}
func (req *Request) HeaderValues(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	result := req.Request.HeaderValues(call.Argument(0).String())
	return r.ToValue(result)
}
func (req *Request) HeaderFields(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	result := req.Request.HeaderFields()
	return r.ToValue(result)
}

func (req *Request) ResponseStatusCode(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(req.Request.ResponseStatusCode())
}
func (req *Request) ResponseBody(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(string(req.Request.ResponseBody()))

}
func (req *Request) ResponseHeader(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(req.Request.ResponseHeader(call.Argument(0).String()))
}
func (req *Request) ResponseHeaderValues(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	result := req.Request.ResponseHeaderValues(call.Argument(0).String())
	return r.ToValue(result)
}
func (req *Request) ResponseHeaderFields(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	result := req.Request.ResponseHeaderFields()
	return r.ToValue(result)
}
func (req *Request) Execute(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	req.Request.MustExecute()
	return nil
}

type Addon struct {
	Addon   *httpaddon.Addon
	Builder Builder
}

func (a *Addon) NewRequest(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	method := call.Argument(0).String()
	url := call.Argument(1).String()
	req := a.Addon.Create(method, url)
	return a.Builder(r, &Request{req})
}
func Create(p herbplugin.Plugin, a *httpaddon.Addon) *Addon {
	return &Addon{
		Addon:   a,
		Builder: DefaultBuilder,
	}
}
