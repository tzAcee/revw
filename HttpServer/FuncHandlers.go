package HttpServer

import (
	"errors"
	"fmt"
	"net/http"
	"revw/Logger"
)

const (
	version = 1
)

type FuncHandlers struct {
	handlers HandlerCollection
}

func NewConcurentFuncHandlers() *FuncHandlers {
	funcHandlers := FuncHandlers{}

	funcHandlers.handlers = CreateRoutes()

	return &funcHandlers
}

func (cfh *FuncHandlers) RegisterUrls() error {
	handleCount := len(cfh.handlers)
	if handleCount == 0 {
		return errors.New("no functions handlers were initialized, so there are no valid routes")
	}
	Logger.GetLogger().Printf("Registering %d URLs\n", handleCount)

	for url, handlers := range cfh.handlers {
		versionedUrl := "/api/v" + fmt.Sprint(version) + url
		localHandlers := handlers

		http.HandleFunc(versionedUrl, func(rw http.ResponseWriter, req *http.Request) {
			handler, ok := localHandlers[req.Method]
			if !ok {
				Logger.GetLogger().Printf("request-method '%v' not allowed on url %v.\n", req.Method, versionedUrl)
				OnHttpErrorWithCode(http.StatusMethodNotAllowed, rw)
				return
			}

			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			jsonStrResponse, handleErr := handler.handle(rw, req)

			if handleErr != nil {
				errMsg := handleErr.Error()
				Logger.GetLogger().Printf("'%v' handle failed on URL '%v', with message '%v'.\n", req.Method, versionedUrl, errMsg)
				OnHttpErrorWithMessage(errMsg, rw)
				return
			}

			fmt.Fprintln(rw, jsonStrResponse)
		})

		Logger.GetLogger().Printf("Registered %d handler on URL '%v'\n", len(handlers), versionedUrl)
	}

	return nil
}
