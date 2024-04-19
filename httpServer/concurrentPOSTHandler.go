package httpServer

import "net/http"

type ConcurrentPOSTHandler struct {
	handleFunc func(http.ResponseWriter, *http.Request)
}

func NewPostHandler(handleFunc func(http.ResponseWriter, *http.Request)) *ConcurrentPOSTHandler {
	return &ConcurrentPOSTHandler{handleFunc}
}

func (ch *ConcurrentPOSTHandler) handle(rw http.ResponseWriter, req *http.Request) {
	ch.handleFunc(rw, req)
}
