package middleware

import (
	"net/http"
	C "../common"
)


var defaultHeaders = map[string]string{
	"Access-Control-Allow-Origin": "*",
	"Cache-Control": "no-cache, private, max-age=0",
	"Pragma": "no-cache",
	"X-Accel-Expires": "0",
	"Content-Type": "application/json",
}


func MiddlewareChain(fn func(C.App, http.ResponseWriter, *http.Request), app C.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Expires", app.Epoch)
		for k, v := range defaultHeaders {
			res.Header().Add(k, v)
		}

		var f func(C.App, http.ResponseWriter, *http.Request) = fn

		f(app, res, req)
	}
} 