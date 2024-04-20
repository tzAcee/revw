package HttpServer

import "net/http"

type Handler interface {
	handle(rw http.ResponseWriter, req *http.Request) error
}
