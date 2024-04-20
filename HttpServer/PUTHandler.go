package HttpServer

import (
	"net/http"
	"revw/HttpServer/HandlerInfo"
)

type PUTHandler struct {
	handlerInfo HandlerInfo.PUTHandlerInfo
	handleFunc  HandlerInfo.PUTHandleFunc
}

func NewPutHandler(handleFunc HandlerInfo.PUTHandleFunc) *PUTHandler {
	return &PUTHandler{HandlerInfo.PUTHandlerInfo{}, handleFunc}
}

func (ch *PUTHandler) handle(rw http.ResponseWriter, req *http.Request) (string, error) {
	return ch.handleFunc(ch.handlerInfo)
}
