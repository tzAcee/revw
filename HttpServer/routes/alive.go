package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"revw/HttpServer/HandlerInfo"
)

type aliveResponse struct {
	Status string
}

func Alive(rw http.ResponseWriter, getHandlerInfo HandlerInfo.GETHandlerInfo) error {
	result, _ := json.Marshal(aliveResponse{"alive"})

	fmt.Fprintln(rw, string(result))

	return errors.New("hello error")
}
