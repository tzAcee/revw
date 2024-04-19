package httpServer

import (
	"fmt"
	"net/http"
)

type HttpServer interface {
	Serve(port uint16) error
}

type ConcurrentHttpServer struct {
	port         uint16
	funcHandlers ConcurentFuncHandlers
}

func NewConcHttpServer() *ConcurrentHttpServer {
	concHttpServer := ConcurrentHttpServer{
		3000,
		*NewConcurentFuncHandlers(),
	}

	concHttpServer.funcHandlers.RegisterUrls()

	return &concHttpServer
}

func (chs *ConcurrentHttpServer) Serve(port uint16) error {
	chs.port = port

	connectionUrl := fmt.Sprintf(":%d", port)
	fmt.Printf("Serving concurrent http-server on '%v'\n", connectionUrl)
	return http.ListenAndServe(connectionUrl, nil)
}
