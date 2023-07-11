package httplua

import (
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/plugins/addons/binaryaddon"
	lua "github.com/yuin/gopher-lua"
)

type Addon struct {
	Addon *binaryaddon.Addon
}

func (a *Addon) Base64Encode(L *lua.LState) int {
	L.Push(lua.LString(a.Addon.Base64Encode([]byte(L.ToString(1)))))
	return 1
}
func (a *Addon) Base64Decode(L *lua.LState) int {
	L.Push(lua.LString(a.Addon.Base64Decode(L.ToString(1))))
	return 1
}
func (a *Addon) Md5Sum(L *lua.LState) int {
	L.Push(lua.LString(a.Addon.Md5Sum([]byte(L.ToString(1)))))
	return 1
}
func (a *Addon) Sha1Sum(L *lua.LState) int {
	L.Push(lua.LString(a.Addon.Sha1Sum([]byte(L.ToString(1)))))
	return 1
}
func (a *Addon) Sha256Sum(L *lua.LState) int {
	L.Push(lua.LString(a.Addon.Sha256Sum([]byte(L.ToString(1)))))
	return 1
}
func (a *Addon) Sha512Sum(L *lua.LState) int {
	L.Push(lua.LString(a.Addon.Sha512Sum([]byte(L.ToString(1)))))
	return 1
}
func (a *Addon) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("Base64Encode", L.NewFunction(a.Base64Encode))
	t.RawSetString("Base64Decode", L.NewFunction(a.Base64Decode))
	t.RawSetString("Md5Sum", L.NewFunction(a.Md5Sum))
	t.RawSetString("Sha1Sum", L.NewFunction(a.Sha1Sum))
	t.RawSetString("Sha256Sum", L.NewFunction(a.Sha256Sum))
	t.RawSetString("Sha512Sum", L.NewFunction(a.Sha512Sum))
	return t
}

func Create(p herbplugin.Plugin) *Addon {
	return &Addon{
		Addon: binaryaddon.Create(p),
	}
}
