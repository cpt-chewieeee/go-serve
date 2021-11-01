package common

import (
	"net/http"
)


func HealthHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"status": "pass"}`))
}