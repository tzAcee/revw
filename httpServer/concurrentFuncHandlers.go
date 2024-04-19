package httpServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type RtError struct {
	ErrorMessage string
}

func OnError(rtErr RtError, rw http.ResponseWriter) {
	errString, err := json.Marshal(rtErr)
	if err != nil {
		fmt.Fprintf(rw, "Could not convert error object, but there was an error with message: %v\n", rtErr.ErrorMessage)
		return
	}

	fmt.Fprintf(rw, "%v\n", string(errString))
}

type handlerCollection map[string]map[string]ConcurrentHandler

func createRoutes() handlerCollection {
	firstCall := true
	handlers := handlerCollection{
		"/": {
			"GET": NewGetHandler(func(rw http.ResponseWriter, req *http.Request) { fmt.Fprintf(rw, "Hello World GET!\n") }),
			"PUT": NewPutHandler(func(rw http.ResponseWriter, req *http.Request) {
				if firstCall {
					firstCall = false
					fmt.Println("Sleeping because of first call")
					time.Sleep(3 * time.Second)
				} else {
					fmt.Println("Not sleeping anymore")
				}

				fmt.Fprintf(rw, "Hello World PUT!\n")
			}),
		},
	}

	return handlers
}

type ConcurentFuncHandlers struct {
	handlers handlerCollection
}

func NewConcurentFuncHandlers() *ConcurentFuncHandlers {
	funcHandlers := ConcurentFuncHandlers{}

	funcHandlers.handlers = createRoutes()

	return &funcHandlers
}

func (cfh *ConcurentFuncHandlers) RegisterUrls() error {
	handleCount := len(cfh.handlers)
	if handleCount == 0 {
		return errors.New("no functions handlers were initialized, so there are no valid routes")
	}
	fmt.Printf("Registering %d URLs\n", handleCount)

	for url, handlers := range cfh.handlers {
		//	var waitGroup sync.WaitGroup

		http.HandleFunc(url, func(rw http.ResponseWriter, req *http.Request) {
			handler, ok := handlers[req.Method]
			if !ok {
				OnError(RtError{fmt.Sprintf("'%v'-request is not supported on URL '%v'\n", req.Method, url)}, rw)
				return
			}

			//	waitGroup.Add(1)
			//	defer waitGroup.Wait()
			//		go func() {
			//			defer waitGroup.Done()
			handler.handle(rw, req)
			//			}()
		})

		fmt.Printf("Registered handleFunc on URL '%v'\n", url)
	}

	return nil
}
