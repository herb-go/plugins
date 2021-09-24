package httpplugin

import (
	"io"
	"net/http"
	"time"
)

type Response struct {
	FinishedAt int64
	StatusCode int
	Header     http.Header
	Body       []byte
}

func ConvertResponse(resp *http.Response) (*Response, error) {
	r := &Response{
		FinishedAt: time.Now().Unix(),
	}
	r.StatusCode = resp.StatusCode
	var err error
	r.Body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r.Header = resp.Header
	return r, nil
}
