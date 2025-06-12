package binaryv8

import (
	"github.com/dop251/goja"
	v8plugin "github.com/herb-go/herbplugin/v8plugin"
	v8 "github.com/herb-go/v8go"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/plugins/addons/binaryaddon"
)

type Addon struct {
	Addon *binaryaddon.Addon
}

func (a *Addon) Base64Encode(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	data := v8plugin.MustGetArg(call, 0)
	if data.IsArrayBuffer() {
		bs, ok := data.SharedArrayBufferGetContents()
		if ok {
			return r.ToValue(a.Addon.Base64Encode(bs.Bytes()))
		}
	}
	return nil
}
func (a *Addon) Base64Decode(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	v8plugin.MustNewGCValue(call.Context(), a.Addon.Base64Decode(v8plugin.MustGetArg(call, 0).String()))
}
func (a *Addon) Md5Sum(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	data := v8plugin.MustGetArg(call, 0)
	if data.IsArrayBuffer() {
		bs, ok := data.(goja.ArrayBuffer)
		if ok {
			return r.ToValue(a.Addon.Md5Sum(bs.Bytes()))
		}
	}
	return nil
}
func (a *Addon) Sha1Sum(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	data := v8plugin.MustGetArg(call, 0)
	if data.IsArrayBuffer() {
		bs, ok := data.(goja.ArrayBuffer)
		if ok {
			return r.ToValue(a.Addon.Sha1Sum(bs.Bytes()))
		}
	}
	return nil
}
func (a *Addon) Sha256Sum(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	data := v8plugin.MustGetArg(call, 0)
	if data.IsArrayBuffer() {
		bs, ok := data.(goja.ArrayBuffer)
		if ok {
			return r.ToValue(a.Addon.Sha256Sum(bs.Bytes()))
		}
	}
	return nil
}
func (a *Addon) Sha512Sum(call *v8.FunctionCallbackInfo) *v8.Value {
	defer call.Release()
	data := v8plugin.MustGetArg(call, 0)
	if data.IsArrayBuffer() {
		bs, ok := data.(goja.ArrayBuffer)
		if ok {
			return r.ToValue(a.Addon.Sha512Sum(bs.Bytes()))
		}
	}
	return nil
}
func (a *Addon) Convert(r *goja.Runtime) *goja.Object {
	obj := r.NewObject()
	obj.Set("Base64Encode", a.Base64Encode)
	obj.Set("Base64Decode", a.Base64Decode)
	obj.Set("Md5Sum", a.Md5Sum)
	obj.Set("Sha1Sum", a.Sha1Sum)
	obj.Set("Sha256Sum", a.Sha256Sum)
	obj.Set("Sha512Sum", a.Sha512Sum)
	return obj
}

func Create(p herbplugin.Plugin) *Addon {
	return &Addon{
		Addon: binaryaddon.Create(p),
	}
}
