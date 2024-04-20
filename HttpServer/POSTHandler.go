package HttpServer

import (
	"net/http"
	"revw/HttpServer/HandlerInfo"
)

type POSTHandler struct {
	handlerInfo HandlerInfo.POSTHandlerInfo
	handleFunc  HandlerInfo.POSTHandleFunc
}

func NewPostHandler(handleFunc HandlerInfo.POSTHandleFunc) *POSTHandler {
	return &POSTHandler{HandlerInfo.POSTHandlerInfo{}, handleFunc}
}

func (ch *POSTHandler) handle(rw http.ResponseWriter, req *http.Request) error {
	return ch.handleFunc(rw, ch.handlerInfo)
}
