package HttpServer

import (
	"net/http"
	"revw/HttpServer/routes"
	"revw/HttpServer/routes/review"
)

type HandlerCollection map[string]map[string]Handler

func CreateRoutes() HandlerCollection {
	handlers := HandlerCollection{
		"/alive": {
			http.MethodGet: NewGetHandler(routes.Alive),
		},
		"/review/request/begin": {
			http.MethodPost: NewPostHandler(review.BeginRequest),
		},
		"/review/read/begin": {
			http.MethodPost: NewPostHandler(review.BeginRead),
		},
		"/review/read/comment/add": {
			http.MethodPost: NewPostHandler(review.CommentAdd),
		},
		"/review/read/comment/delete": {
			http.MethodPost: NewPostHandler(review.CommentDelete),
		},
		"/review/read/comment/edit": {
			http.MethodPost: NewPostHandler(review.CommentEdit),
		},
	}

	return handlers
}
