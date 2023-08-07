package httplua

import (
	"net/url"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/plugins/addons/httpaddon"
	lua "github.com/yuin/gopher-lua"
)

type Builder func(L *lua.LState, req *Request) *lua.LTable

var DefaultBuilder = func(L *lua.LState, req *Request) *lua.LTable {
	t := L.NewTable()
	t.RawSetString("GetID", L.NewFunction(req.GetID))
	t.RawSetString("GetURL", L.NewFunction(req.GetURL))
	t.RawSetString("SetURL", L.NewFunction(req.SetURL))
	t.RawSetString("GetMethod", L.NewFunction(req.GetMethod))
	t.RawSetString("SetMethod", L.NewFunction(req.SetMethod))
	t.RawSetString("GetBody", L.NewFunction(req.GetBody))
	t.RawSetString("SetBody", L.NewFunction(req.SetBody))
	t.RawSetString("FinishedAt", L.NewFunction(req.FinishedAt))
	t.RawSetString("ExecuteStatus", L.NewFunction(req.ExecuteStatus))
	t.RawSetString("ResetHeader", L.NewFunction(req.ResetHeader))
	t.RawSetString("SetHeader", L.NewFunction(req.SetHeader))
	t.RawSetString("AddHeader", L.NewFunction(req.AddHeader))
	t.RawSetString("DelHeader", L.NewFunction(req.DelHeader))
	t.RawSetString("GetHeader", L.NewFunction(req.GetHeader))
	t.RawSetString("HeaderValues", L.NewFunction(req.HeaderValues))
	t.RawSetString("HeaderFields", L.NewFunction(req.HeaderFields))
	t.RawSetString("ResponseStatusCode", L.NewFunction(req.ResponseStatusCode))
	t.RawSetString("ResponseBody", L.NewFunction(req.ResponseBody))
	t.RawSetString("ResponseHeader", L.NewFunction(req.ResponseHeader))
	t.RawSetString("ResponseHeaderValues", L.NewFunction(req.ResponseHeaderValues))
	t.RawSetString("ResponseHeaderFields", L.NewFunction(req.ResponseHeaderFields))
	t.RawSetString("Execute", L.NewFunction(req.Execute))
	return t
}

type Request struct {
	Request *httpaddon.Request
}

func (req *Request) PushStrings(L *lua.LState, data []string) {
	t := L.NewTable()
	for _, v := range data {
		t.Append(lua.LString(v))
	}
	L.Push(t)
}
func (req *Request) GetID(L *lua.LState) int {
	L.Push(lua.LString(req.Request.GetID()))
	return 1
}

func (req *Request) GetURL(L *lua.LState) int {
	L.Push(lua.LString(req.Request.GetURL()))
	return 1
}
func (req *Request) SetURL(L *lua.LState) int {
	req.Request.SetURL(L.ToString((1)))
	return 0
}
func (req *Request) GetProxy(L *lua.LState) int {
	L.Push(lua.LString(req.Request.GetProxy()))
	return 1
}
func (req *Request) SetProxy(L *lua.LState) int {
	req.Request.SetProxy(L.ToString((1)))
	return 0
}
func (req *Request) GetMethod(L *lua.LState) int {
	L.Push(lua.LString(req.Request.GetMethod()))
	return 1
}
func (req *Request) SetMethod(L *lua.LState) int {
	req.Request.SetMethod(L.ToString(1))
	return 0

}
func (req *Request) GetBody(L *lua.LState) int {
	L.Push(lua.LString(req.Request.GetBody()))
	return 1
}
func (req *Request) SetBody(L *lua.LState) int {
	req.Request.SetBody([]byte(L.ToString(1)))
	return 0

}
func (req *Request) FinishedAt(L *lua.LState) int {
	L.Push(lua.LNumber(req.Request.FinishedAt()))
	return 1

}
func (req *Request) ExecuteStatus(L *lua.LState) int {
	L.Push(lua.LNumber(req.Request.ExecuteStauts()))
	return 1
}
func (req *Request) ResetHeader(L *lua.LState) int {
	req.Request.ResetHeader()
	return 0
}
func (req *Request) SetHeader(L *lua.LState) int {
	req.Request.SetHeader(L.ToString(1), L.ToString(2))
	return 0
}
func (req *Request) AddHeader(L *lua.LState) int {
	req.Request.AddHeader(L.ToString(1), L.ToString(2))
	return 0
}
func (req *Request) DelHeader(L *lua.LState) int {
	req.Request.DelHeader(L.ToString(1))
	return 0

}
func (req *Request) GetHeader(L *lua.LState) int {
	L.Push(lua.LString(req.Request.GetHeader(L.ToString(1))))
	return 1

}
func (req *Request) HeaderValues(L *lua.LState) int {
	result := req.Request.HeaderValues(L.ToString(1))
	req.PushStrings(L, result)
	return 1
}
func (req *Request) HeaderFields(L *lua.LState) int {
	result := req.Request.HeaderFields()
	req.PushStrings(L, result)
	return 1
}

func (req *Request) ResponseStatusCode(L *lua.LState) int {
	L.Push(lua.LNumber((req.Request.ResponseStatusCode())))
	return 1
}
func (req *Request) ResponseBody(L *lua.LState) int {
	L.Push(lua.LString((req.Request.ResponseBody())))
	return 1

}
func (req *Request) ResponseHeader(L *lua.LState) int {
	L.Push(lua.LString(req.Request.ResponseHeader(L.ToString(1))))
	return 1
}
func (req *Request) ResponseHeaderValues(L *lua.LState) int {
	result := req.Request.ResponseHeaderValues(L.ToString(1))
	req.PushStrings(L, result)
	return 1
}
func (req *Request) ResponseHeaderFields(L *lua.LState) int {
	result := req.Request.ResponseHeaderFields()
	req.PushStrings(L, result)
	return 1
}
func (req *Request) Execute(L *lua.LState) int {
	req.Request.MustExecute()
	return 0
}

type Addon struct {
	Addon   *httpaddon.Addon
	Builder Builder
}

func (a *Addon) ParseURL(L *lua.LState) int {
	rawurl := L.ToString(1)
	u, err := url.Parse(rawurl)
	if err != nil {
		L.Push(lua.LNil)
		return 1
	}
	result := L.NewTable()
	result.RawSetString("Host", lua.LString(u.Host))
	result.RawSetString("Hostname", lua.LString(u.Host))
	result.RawSetString("Scheme", lua.LString(u.Scheme))
	result.RawSetString("Path", lua.LString(u.Path))
	result.RawSetString("Query", lua.LString(u.RawQuery))
	result.RawSetString("User", lua.LString(u.User.Username()))
	p, _ := u.User.Password()
	result.RawSetString("Password", lua.LString(p))
	result.RawSetString("Port", lua.LString(u.Port()))
	result.RawSetString("Fragment", lua.LString(u.Fragment))
	L.Push(result)
	return 1
}
func (a *Addon) NewRequest(L *lua.LState) int {
	method := L.ToString(1)
	url := L.ToString(2)
	req := a.Addon.Create(method, url)
	L.Push(a.Builder(L, &Request{req}))
	return 1
}

func (a *Addon) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("New", L.NewFunction(a.NewRequest))
	t.RawSetString("ParseURL", L.NewFunction(a.ParseURL))
	return t
}
func Create(p herbplugin.Plugin) *Addon {
	return &Addon{
		Addon:   httpaddon.Create(p),
		Builder: DefaultBuilder,
	}
}
