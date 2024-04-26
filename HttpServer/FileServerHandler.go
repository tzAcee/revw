package HttpServer

import (
	"net/http"
	"revw/HttpServer/HandlerInfo"
)

type FileServerHandler struct {
	handlerInfo   HandlerInfo.FileServerHandleInfo
	fileServerDir string
}

func NewFSHandler(fileDir string) *FileServerHandler {
	return &FileServerHandler{HandlerInfo.FileServerHandleInfo{}, fileDir}
}

func (ch *FileServerHandler) handle(rw http.ResponseWriter, req *http.Request) (string, error) {
	return ch.fileServerDir, nil
}
