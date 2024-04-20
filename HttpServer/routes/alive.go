package routes

import (
	"encoding/json"
	"revw/HttpServer/HandlerInfo"
)

type aliveResponse struct {
	Status string
}

func Alive(getHandlerInfo HandlerInfo.GETHandlerInfo) (string, error) {
	result, _ := json.Marshal(aliveResponse{"alive"})

	return string(result), nil

	//return errors.New("hello error")
}
