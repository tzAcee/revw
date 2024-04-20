package main

import (
	"flag"
	"revw/HttpServer"
)

func main() {
	var cfg HttpServer.Config

	flag.UintVar(&cfg.Port, "port", 3000, "server port of the API")
	flag.Parse()

	var httpServer HttpServer.HttpServerBase = HttpServer.NewHttpServer()

	httpServer.Serve(cfg)
}
