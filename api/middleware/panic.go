package middleware

import (
	"go-blog/api"
	"log"
	"net/http"
	"runtime/debug"
)

type RecoverHandler struct {
	api.IErrorHandler
}

func NewRecoverHandler(handler api.IErrorHandler) RecoverHandler{
	return RecoverHandler{
		IErrorHandler: handler,
	}
}

// AccessLogger is a handler decorator which logs the method, path, duration, and
// status for every request
func (rH RecoverHandler)RecoverFromPanic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				m := r.Method
				p := r.URL.Path
				log.Printf("%v %v panic(%v)", m, p, rec)
				log.Printf("stacktrace from panic: \n%v" + string(debug.Stack()))
				rH.InternalServerError(w, r)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
