package middleware

import (
	"go-blog/api"
	"log"
	"net/http"
	"runtime/debug"
)

//RecoverHandler recovers from panic
type RecoverHandler struct {
	api.IErrorHandler
}

//NewRecoverHandler returns new
func NewRecoverHandler(handler api.IErrorHandler) RecoverHandler {
	return RecoverHandler{
		IErrorHandler: handler,
	}
}

//RecoverFromPanic turns panics into 500s
func (rH RecoverHandler) RecoverFromPanic(h http.Handler) http.Handler {
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
