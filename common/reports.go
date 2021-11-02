package common

import (
	"net/http"
	"strconv"
	"runtime"
)


func HealthHandler(a App, res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	res.Write([]byte(`{"status": "pass", "port": ` + strconv.Itoa(a.Port) + 
		`, "path": "` + a.Path + 
		`", "expires": "` + a.Epoch + 
		` MiB", "alloc": "` + strconv.FormatUint(bToMb(m.Alloc), 10) +
		` MiB", "total_alloc": "` + strconv.FormatUint(bToMb(m.TotalAlloc), 10) +
		` MiB", "sys": "` + strconv.FormatUint(bToMb(m.Sys), 10) +
		`", "objects": "` + strconv.FormatUint(m.HeapObjects, 10) +
		`", "numgc": "` + strconv.Itoa(int(m.NumGC)) +
		`"}`))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}