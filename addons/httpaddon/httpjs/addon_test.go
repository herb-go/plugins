package httpjs

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/herb-go/plugins/addons/httpaddon"

	"github.com/dop251/goja"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
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
	i := jsplugin.NewInitializer()
	i.Entry = "test.js"
	module := herbplugin.CreateModule(
		"test",
		func(ctx context.Context, p herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			plugin := p.(*jsplugin.Plugin)
			plugin.Runtime.Set("HTTP", Create(p).Convert(plugin.Runtime))
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
	p := jsplugin.MustCreatePlugin(i)
	herbplugin.Lanuch(p, opt)
	fn, ok := goja.AssertFunction(p.Runtime.Get("test"))
	if !ok {
		t.Fatal()
	}
	_, err = fn(goja.Undefined(), p.Runtime.ToValue(s.URL))
	if err != nil {
		panic(err)
	}
}
