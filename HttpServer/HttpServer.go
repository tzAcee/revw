package HttpServer

import (
	"fmt"
	"net/http"
	"revw/Logger"
)

type Config struct {
	Port uint
}

type HttpServerBase interface {
	Serve(cfg Config) error
}

type HttpServer struct {
	funcHandlers FuncHandlers
}

func NewHttpServer() *HttpServer {
	concHttpServer := HttpServer{
		*NewConcurentFuncHandlers(),
	}

	err := concHttpServer.funcHandlers.RegisterUrls()
	if err != nil {
		Logger.GetLogger().Fatal(err)
		return nil
	}

	return &concHttpServer
}

func (chs *HttpServer) Serve(cfg Config) error {
	connectionUrl := fmt.Sprintf(":%d", cfg.Port)
	Logger.GetLogger().Printf("Serving http-server on '%v'\n", connectionUrl)
	return http.ListenAndServe(connectionUrl, nil)
}
