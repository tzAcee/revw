package httpServer

import "net/http"

type ConcurrentPUTHandler struct {
	handleFunc func(http.ResponseWriter, *http.Request)
}

func NewPutHandler(handleFunc func(http.ResponseWriter, *http.Request)) *ConcurrentPUTHandler {
	return &ConcurrentPUTHandler{handleFunc}
}

func (ch *ConcurrentPUTHandler) handle(rw http.ResponseWriter, req *http.Request) {
	ch.handleFunc(rw, req)
}
