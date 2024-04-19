package main

import "revw/httpServer"

func main() {
	var httpServer httpServer.HttpServer = httpServer.NewConcHttpServer()

	httpServer.Serve(3500)
}
