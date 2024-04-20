package HttpServer

import (
	"net/http"
	"revw/HttpServer/HandlerInfo"
)

type GETHandler struct {
	handlerInfo HandlerInfo.GETHandlerInfo
	handleFunc  HandlerInfo.GETHandleFunc
}

func NewGetHandler(handleFunc HandlerInfo.GETHandleFunc) *GETHandler {
	return &GETHandler{HandlerInfo.GETHandlerInfo{}, handleFunc}
}

func (ch *GETHandler) handle(rw http.ResponseWriter, req *http.Request) error {
	return ch.handleFunc(rw, ch.handlerInfo)
}
