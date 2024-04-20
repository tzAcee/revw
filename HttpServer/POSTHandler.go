package HttpServer

import (
	"encoding/json"
	"net/http"
	"revw/HttpServer/HandlerInfo"
)

type POSTHandler struct {
	handlerInfo HandlerInfo.POSTHandlerInfo
	handleFunc  HandlerInfo.POSTHandleFunc
}

func NewPostHandler(handleFunc HandlerInfo.POSTHandleFunc) *POSTHandler {
	return &POSTHandler{HandlerInfo.POSTHandlerInfo{}, handleFunc}
}

func (ch *POSTHandler) handle(rw http.ResponseWriter, req *http.Request) (string, error) {
	if req.Body != http.NoBody {
		err := json.NewDecoder(req.Body).Decode(&ch.handlerInfo.RequestBody)

		if err != nil {
			return "", err
		}
	}

	return ch.handleFunc(ch.handlerInfo)
}
