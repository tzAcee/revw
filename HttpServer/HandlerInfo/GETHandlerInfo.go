package HandlerInfo

import "net/http"

type GETHandlerInfo struct {
	Params map[string]any
}

type GETHandleFunc func(http.ResponseWriter, GETHandlerInfo) error
