package httpaddon

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/herb-go/herbplugin"
)

func echoAction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	for key, value := range r.Header {
		for _, v := range value {
			w.Header().Add("Echo-"+key, v)
		}
	}
	w.Write(body)
}
func catch(h func()) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			e, ok := r.(error)
			if ok {
				err = e
			}
		}
	}()
	h()
	return
}
func TestPermission(t *testing.T) {
	app := &http.ServeMux{}
	app.HandleFunc("/echo", echoAction)
	server := httptest.NewServer(app)
	defer server.Close()
	opt := herbplugin.NewOptions()
	u, err := url.Parse(server.URL)
	if err != nil {
		panic(err)
	}

	opt.Trusted.Domains = append(opt.Trusted.Domains, u.Host, "256.256.256.256")
	plugin := herbplugin.New()
	plugin.SetPluginOptions(opt)
	addon := Create(plugin)
	req := addon.Create("GET", server.URL)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	req.SetURL(server.URL + "/echo")
	req.AsyncExecute(func(err error) {
		if !herbplugin.IsUnauthorizePermissionError(err) {
			panic(err)
		}
		wg.Done()
	})
	wg.Wait()
	if req.Status != StatusReady {
		t.Fatal(req)
	}
}

func TestAddon(t *testing.T) {
	var err error
	app := &http.ServeMux{}
	app.HandleFunc("/echo", echoAction)
	server := httptest.NewServer(app)
	defer server.Close()
	opt := herbplugin.NewOptions()
	opt.Permissions = append(opt.Permissions, Permission)
	u, err := url.Parse(server.URL)
	if err != nil {
		panic(err)
	}
	opt.Trusted.Domains = append(opt.Trusted.Domains, u.Host, "256.256.256.256")
	plugin := herbplugin.New()
	plugin.SetPluginOptions(opt)
	addon := Create(plugin)
	req := addon.Create("GET", server.URL)
	err = catch(func() {
		req.ResponseStatusCode()
	})
	if err != ErrRequestNotExecuted {
		t.Fatal(err)
	}
	err = catch(func() {
		req.ResponseHeader("test")
	})
	if err != ErrRequestNotExecuted {
		t.Fatal(err)
	}
	err = catch(func() {
		req.ResponseHeaderFields()
	})
	if err != ErrRequestNotExecuted {
		t.Fatal(err)
	}
	err = catch(func() {
		req.ResponseHeaderFields()
	})
	if err != ErrRequestNotExecuted {
		t.Fatal(err)
	}
	err = catch(func() {
		req.ResponseStatusCode()
	})
	if err != ErrRequestNotExecuted {
		t.Fatal(err)
	}
	err = catch(func() {
		req.ResponseHeaderValues("test")
	})
	if err != ErrRequestNotExecuted {
		t.Fatal(err)
	}
	err = catch(func() {
		req.FinishedAt()
	})
	if err != ErrRequestNotExecuted {
		t.Fatal(err)
	}
	err = catch(func() {
		req.ResponseBody()
	})
	if err != ErrRequestNotExecuted {
		t.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	req.SetHeader("Test1", "testvalue1")
	req.AddHeader("Test1", "testvalueb")
	req.SetHeader("Test2", "testvalue2")
	req.SetBody([]byte("testbody"))
	req.SetMethod("POST")
	req.SetURL(server.URL + "/echo")
	req.AsyncExecute(func(err error) {
		if err != nil {
			panic(err)
		}
		wg.Done()
	})
	wg.Wait()

	if req.FinishedAt() == 0 {
		t.Fatal(req)
	}
	if req.Status != StatusSuccess {
		t.Fatal(req)
	}
	if len(req.ResponseHeaderFields()) < 2 {
		t.Fatal(req)
	}
	if req.ResponseHeader("Echo-Test2") != "testvalue2" {
		t.Fatal(req)
	}
	if strings.Join(req.ResponseHeaderValues("Echo-Test1"), " ") != "testvalue1 testvalueb" {
		t.Fatal(req)
	}
	if string(req.ResponseBody()) != "testbody" {
		t.Fatal(req)
	}
	if req.ResponseStatusCode() != 200 {
		t.Fatal(req)
	}
	wg = &sync.WaitGroup{}
	wg.Add(1)
	req.AsyncExecute(func(err error) {
		if err != ErrRequestExecuted {
			panic(err)
		}
		wg.Done()
	})
	wg.Wait()
	wg = &sync.WaitGroup{}
	wg.Add(1)
	req = addon.Create("GET", "WRONGSCHEME://256.256.256.256/")
	req.AsyncExecute(func(err error) {
		if err == nil {
			panic(err)
		}
		wg.Done()
	})
	wg.Wait()
	if req.Status != StatusFail {
		t.Fatal(req)
	}
	wg = &sync.WaitGroup{}
	wg.Add(1)
	req = addon.Create("GET", "http://wwww/")
	req.AsyncExecute(func(err error) {
		if !herbplugin.IsUnauthorizeDomainError(err) {
			panic(err)
		}
		wg.Done()
	})
	wg.Wait()
	if req.Status != StatusReady {
		t.Fatal(req)
	}
	wg = &sync.WaitGroup{}
	wg.Add(1)
	req = addon.Create("GET", "http://256.256.256.256/")
	req.SetHeader("host", "abc")
	req.AsyncExecute(func(err error) {
		if !herbplugin.IsUnauthorizeDomainError(err) {
			panic(err)
		}
		wg.Done()
	})
	wg.Wait()
	if req.Status != StatusReady {
		t.Fatal(req)
	}
}
func TestRequest(t *testing.T) {
	var data []byte
	var fields []string
	var value string
	var values []string
	var status int
	plugin := herbplugin.New()
	addon := Create(plugin)
	req := addon.Create("GET", "")
	if req.GetID() == "" {
		t.Fatal(req)
	}
	if req.GetMethod() != "GET" {
		t.Fatal(req)
	}
	req.SetMethod("POST")
	if req.GetMethod() != "POST" {
		t.Fatal(req)
	}
	if req.GetURL() != "" {
		t.Fatal(req)
	}
	url := "/echo"
	req.SetURL(url)
	if req.GetURL() != url {
		t.Fatal(req)
	}
	data = req.GetBody()
	if len(data) != 0 {
		t.Fatal(req)
	}
	req.SetBody([]byte("testbody"))
	data = req.GetBody()
	if string(data) != "testbody" {
		t.Fatal(req)
	}
	fields = req.HeaderFields()
	if len(fields) != 0 {
		t.Fatal(req)
	}
	req.SetHeader("test", "testvalue")
	fields = req.HeaderFields()
	if len(fields) != 1 || fields[0] != "Test" {
		t.Fatal(req)
	}
	values = req.HeaderValues("test")
	if len(values) != 1 || values[0] != "testvalue" {
		t.Fatal(req)
	}
	req.AddHeader("test", "testvalue2")
	value = req.GetHeader("test")
	if value != "testvalue" {
		t.Fatal(req)
	}
	values = req.HeaderValues("test")
	if len(values) != 2 || values[0] != "testvalue" || values[1] != "testvalue2" {
		t.Fatal(req)
	}

	req.SetHeader("tes", "tesvalue")
	req.SetHeader("test2", "test2value")
	fields = req.HeaderFields()
	if len(fields) != 3 || fields[0] != "Tes" && fields[1] != "Test" && fields[2] != "Test2" {
		t.Fatal(req)
	}
	req.DelHeader("tes")
	fields = req.HeaderFields()
	if len(fields) != 2 || fields[0] != "Test" && fields[1] != "Test2" {
		t.Fatal(req)
	}
	req.ResetHeader()
	fields = req.HeaderFields()
	if len(fields) != 0 {
		t.Fatal(req)
	}
	status = req.ExecuteStauts()
	if status != StatusReady {
		t.Fatal(req)
	}
}
