package HttpServer

import "net/http"

func OnHttpErrorWithCode(code int, rw http.ResponseWriter) {
	http.Error(rw, http.StatusText(code), code)
}

func OnHttpErrorWithMessage(msg string, rw http.ResponseWriter) {
	http.Error(rw, msg, http.StatusInternalServerError)
}
