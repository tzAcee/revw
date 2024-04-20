package HandlerInfo

import "net/http"

type PUTHandlerInfo struct {
	Params map[string]any
}

type PUTHandleFunc func(http.ResponseWriter, PUTHandlerInfo) error
