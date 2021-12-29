package httplua

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/herb-go/plugins/addons/httpaddon"
	lua "github.com/yuin/gopher-lua"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
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
	i := lua51plugin.NewInitializer()
	i.Entry = "test.lua"
	module := herbplugin.CreateModule(
		"test",
		func(ctx context.Context, p herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			plugin := p.(*lua51plugin.Plugin)
			plugin.LState.SetGlobal("HTTP", Create(p).Convert(plugin.LState))
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
	p := lua51plugin.MustCreatePlugin(i)
	herbplugin.Lanuch(p, opt)
	fn := p.LState.GetGlobal("test")
	if err := p.LState.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, lua.LString(s.URL)); err != nil {
		t.Fatal(err)
	}

}
