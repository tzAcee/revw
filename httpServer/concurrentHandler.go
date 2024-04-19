package httpServer

import "net/http"

type ConcurrentHandler interface {
	handle(rw http.ResponseWriter, req *http.Request)
}
