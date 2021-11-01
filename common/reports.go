package common

import (
	"net/http"
	"strconv"
)


func HealthHandler(a App, res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"status": "pass", "port": ` + strconv.Itoa(a.Port) + `, "path": "` + a.Path + `", "expires": "` + a.Epoch + `"}`))
}