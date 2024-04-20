package HttpServer

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port uint
}

type HttpServerBase interface {
	Serve(cfg Config) error
}

type HttpServer struct {
	logger       *log.Logger
	funcHandlers FuncHandlers
}

func NewHttpServer() *HttpServer {
	logger := log.New(os.Stdout, "[revw-log] ", log.Ldate|log.Ltime)
	concHttpServer := HttpServer{
		logger,
		*NewConcurentFuncHandlers(logger),
	}

	err := concHttpServer.funcHandlers.RegisterUrls()
	if err != nil {
		logger.Fatal(err)
		return nil
	}

	return &concHttpServer
}

func (chs *HttpServer) Serve(cfg Config) error {
	connectionUrl := fmt.Sprintf(":%d", cfg.Port)
	chs.logger.Printf("Serving http-server on '%v'\n", connectionUrl)
	return http.ListenAndServe(connectionUrl, nil)
}
