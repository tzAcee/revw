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

type CommentRequest struct {
	Index uint
	Text  string
}

type BeginReadResponse BeginRequestResponse
type AddCommentResponse BeginReadResponse

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

// expects {"ReviewRequestID": "reviewRequestID", "ReaderID": "readerID", "CommentIndex": 123123, "CommentText": "comment"}
func CommentAdd(handlerInfo HandlerInfo.POSTHandlerInfo) (string, error) {
	reviewRequestID, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "ReviewRequestID")
	if err != nil {
		return "", err
	}
	if len(reviewRequestID) == 0 {
		return "", errors.New("ReviewRequestID is empty")
	}

	readerID, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "ReaderID")
	if err != nil {
		return "", err
	}
	if len(readerID) == 0 {
		return "", errors.New("ReaderID is empty")
	}

	commentText, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "CommentText")
	if err != nil {
		return "", err
	}
	if len(commentText) == 0 {
		return "", errors.New("CommentText is empty")
	}
	commentIndex, err := HandlerInfo.GetEntryOfRequestbody[float64](handlerInfo.RequestBody, "CommentIndex")
	if err != nil {
		return "", err
	}

	reviewSession, ok := RevwBL.GetSessionsManager().GetSessionById(reviewRequestID)
	if reviewSession == nil || !ok {
		return "", errors.New("could not find review session")
	}

	reader, ok := reviewSession.GetReaderByID(readerID)
	if !ok {
		return "", errors.New("could not find a reader with this id")
	}

	comment := reader.AddComment(commentText, uint(commentIndex))

	response, _ := json.Marshal(AddCommentResponse{comment.ID})

	return string(response), nil
}

// expects {"ReviewRequestID": "reviewRequestID", "ReaderID": "readerID", "CommentID": "commentID"}
func CommentDelete(handlerInfo HandlerInfo.POSTHandlerInfo) (string, error) {
	reviewRequestID, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "ReviewRequestID")
	if err != nil {
		return "", err
	}
	if len(reviewRequestID) == 0 {
		return "", errors.New("ReviewRequestID is empty")
	}

	readerID, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "ReaderID")
	if err != nil {
		return "", err
	}
	if len(readerID) == 0 {
		return "", errors.New("ReaderID is empty")
	}

	commentID, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "CommentID")
	if err != nil {
		return "", err
	}
	if len(commentID) == 0 {
		return "", errors.New("CommentID is empty")
	}

	reviewSession, ok := RevwBL.GetSessionsManager().GetSessionById(reviewRequestID)
	if reviewSession == nil || !ok {
		return "", errors.New("could not find review session")
	}

	reader, ok := reviewSession.GetReaderByID(readerID)
	if !ok {
		return "", errors.New("could not find a reader with this id")
	}

	reader.DeleteComment(commentID)

	response, _ := json.Marshal(Result{true})

	return string(response), nil
}

// expects {"ReviewRequestID": "reviewRequestID", "ReaderID": "readerID", "CommentID": "commentID", "CommentText": "newCommentText"}
func CommentEdit(handlerInfo HandlerInfo.POSTHandlerInfo) (string, error) {
	reviewRequestID, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "ReviewRequestID")
	if err != nil {
		return "", err
	}
	if len(reviewRequestID) == 0 {
		return "", errors.New("ReviewRequestID is empty")
	}

	readerID, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "ReaderID")
	if err != nil {
		return "", err
	}
	if len(readerID) == 0 {
		return "", errors.New("ReaderID is empty")
	}

	commentID, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "CommentID")
	if err != nil {
		return "", err
	}
	if len(commentID) == 0 {
		return "", errors.New("CommentID is empty")
	}

	commentText, err := HandlerInfo.GetEntryOfRequestbody[string](handlerInfo.RequestBody, "CommentText")
	if err != nil {
		return "", err
	}
	if len(commentText) == 0 {
		return "", errors.New("CommentID is empty")
	}

	reviewSession, ok := RevwBL.GetSessionsManager().GetSessionById(reviewRequestID)
	if reviewSession == nil || !ok {
		return "", errors.New("could not find review session")
	}

	reader, ok := reviewSession.GetReaderByID(readerID)
	if !ok {
		return "", errors.New("could not find a reader with this id")
	}

	editResult := reader.EditComment(commentID, commentText)

	if editResult != nil {
		return "", editResult
	}

	response, _ := json.Marshal(Result{true})

	return string(response), nil
}
