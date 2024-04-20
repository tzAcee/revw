package HttpServer

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	version = 1
)

type FuncHandlers struct {
	logger   *log.Logger
	handlers HandlerCollection
}

func NewConcurentFuncHandlers(logger *log.Logger) *FuncHandlers {
	funcHandlers := FuncHandlers{}

	funcHandlers.logger = logger
	funcHandlers.handlers = CreateRoutes()

	return &funcHandlers
}

func (cfh *FuncHandlers) RegisterUrls() error {
	handleCount := len(cfh.handlers)
	if handleCount == 0 {
		return errors.New("no functions handlers were initialized, so there are no valid routes")
	}
	cfh.logger.Printf("Registering %d URLs\n", handleCount)

	for url, handlers := range cfh.handlers {
		versionedUrl := "/v" + fmt.Sprint(version) + url

		http.HandleFunc(versionedUrl, func(rw http.ResponseWriter, req *http.Request) {
			handler, ok := handlers[req.Method]
			if !ok {
				cfh.logger.Printf("request-method '%v' not allowed on url %v.\n", req.Method, versionedUrl)
				OnHttpErrorWithCode(http.StatusMethodNotAllowed, rw)
				return
			}

			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			handleErr := handler.handle(rw, req)

			if handleErr != nil {
				errMsg := handleErr.Error()
				cfh.logger.Printf("'%v' handle failed on URL '%v', with message '%v'.\n", req.Method, versionedUrl, errMsg)
				OnHttpErrorWithMessage(errMsg, rw)
			}
		})

		cfh.logger.Printf("Registered handleFunc on URL '%v'\n", versionedUrl)
	}

	return nil
}
