package httpServer

import "net/http"

type ConcurrentGETHandler struct {
	handleFunc func(http.ResponseWriter, *http.Request)
}

func NewGetHandler(handleFunc func(http.ResponseWriter, *http.Request)) *ConcurrentGETHandler {
	return &ConcurrentGETHandler{handleFunc}
}

func (ch *ConcurrentGETHandler) handle(rw http.ResponseWriter, req *http.Request) {
	ch.handleFunc(rw, req)
}
