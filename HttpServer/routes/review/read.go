package review

import (
	"encoding/json"
	"errors"
	"revw/HttpServer/HandlerInfo"
	"revw/RevwBL"
)

type Result struct {
	Success bool
}

type BeginReadResponse BeginRequestResponse

// expects {"ID": "reviewRequestID"}
func BeginRead(handlerInfo HandlerInfo.POSTHandlerInfo) (string, error) {
	id, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "ID")
	if err != nil {
		return "", err
	}

	reviewSession, ok := RevwBL.GetSessionsManager().GetSessionById(id)
	if reviewSession == nil || !ok {
		return "", errors.New("could not find review session")
	}

	newReader := reviewSession.AddReader()
	response, _ := json.Marshal(BeginReadResponse{newReader.ID})

	return string(response), nil
}
