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

// expects {"ReviewID": "id"}
func GetReview(handlerInfo HandlerInfo.POSTHandlerInfo) (string, error) {
	reviewID, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "ReviewID")
	if err != nil {
		return "", err
	}

	if len(reviewID) == 0 {
		return "", errors.New("review id cannot be empty")
	}

	reviewSession, ok := RevwBL.GetSessionsManager().GetSessionById(reviewID)
	if !ok {
		return "", errors.New("could not find review session")
	}

	response, _ := json.Marshal(reviewSession)

	return string(response), nil
}
