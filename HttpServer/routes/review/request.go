package review

import (
	"encoding/json"
	"errors"
	"revw/HttpServer/HandlerInfo"
	"revw/RevwBL"
)

type BeginRequestResponse struct {
	ID string
}

// expects {"Text": "someReviewText"}
func BeginRequest(handlerInfo HandlerInfo.POSTHandlerInfo) (string, error) {
	reviewText, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "Text")
	if err != nil {
		return "", err
	}

	if len(reviewText) == 0 {
		return "", errors.New("review text cannot be empty")
	}

	reviewSession := RevwBL.GetSessionsManager().CreateReviewSession(&reviewText)
	if reviewSession == nil {
		return "", errors.New("could not create review session")
	}

	response, _ := json.Marshal(BeginRequestResponse{reviewSession.ID})

	return string(response), nil
}
