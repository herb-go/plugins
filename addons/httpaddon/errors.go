package httpaddon

import "errors"

var ErrRequestNotExecuted = errors.New("httpplugin: request not executed")
var ErrRequestExecuted = errors.New("httpplugin: request executed")
