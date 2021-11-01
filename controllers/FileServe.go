package controllers 

import (
	"net/http"
	C "../common"
)

func FileHttpServer(a C.App, w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(a.Path)).ServeHTTP(w, r)
}