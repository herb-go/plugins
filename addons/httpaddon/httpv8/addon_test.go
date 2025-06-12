package httpv8

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/herb-go/plugins/addons/httpaddon"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/v8plugin"
	v8 "github.com/herb-go/v8go"
)

func TestAddon(t *testing.T) {
	app := &http.ServeMux{}
	s := httptest.NewServer(app)
	defer s.Close()
	u, err := url.Parse(s.URL)
	if err != nil {
		panic(err)
	}
	opt := herbplugin.NewOptions()
	opt.Permissions = append(opt.Permissions, httpaddon.Permission)
	opt.Trusted.Domains = append(opt.Trusted.Domains, u.Host)
	opt.GetLocation().Path = "."
	i := v8plugin.NewInitializer()
	i.Entry = "test.js"
	module := herbplugin.CreateModule(
		"test",
		func(ctx context.Context, p herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			plugin := p.(*v8plugin.Plugin)
			plugin.Runtime.Global().Set("HTTP", Create(p).Convert(plugin.Runtime))
			next(ctx, p)
		},
		func(ctx context.Context, p herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			next(ctx, p)
		},
		func(ctx context.Context, p herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			next(ctx, p)
		},
	)
	i.Modules = append(i.Modules, module)
	p := v8plugin.MustCreatePlugin(i)
	herbplugin.Lanuch(p, opt)
	test, err := p.Runtime.Global().Get("test")
	if err != nil {
		t.Fatal(err)
	}
	fn, err := test.AsFunction()
	if err != nil {
		t.Fatal(err)
	}
	_, err = fn.Call(v8.Undefined(p.Runtime.Isolate()), v8plugin.MustNewGCValue(p.Runtime, s.URL))
	if err != nil {
		panic(err)
	}
}
