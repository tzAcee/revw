package HandlerInfo

import "net/http"

type POSTHandlerInfo struct {
	Params map[string]any
}

type POSTHandleFunc func(http.ResponseWriter, POSTHandlerInfo) error
