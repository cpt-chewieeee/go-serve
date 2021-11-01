package main

import (
	"flag"
	// "net/http"
	// "strconv"
	"time"
	C "./common"
	"fmt"
)


var (
	port = flag.Int("port", 8000, "port number")
	path = flag.String("path", ".", "path")
	epoch = time.Unix(0, 0).Format(time.RFC1123)
)

// var defaultHeaders = map[string]string{
// 	"Access-Control-Allow-Origin": "*",
// 	"Expires": epoch,
// 	"Cache-Control": "no-cache, private, max-age=0",
// 	"Pragma": "no-cache",
// 	"X-Accel-Expires": "0",
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	for k, v := range defaultHeaders {
// 		w.Header().Add(k, v)
// 	}

// 	http.FileServer(http.Dir(*path)).ServeHTTP(w, r)
// }

func newApp () *C.App {
	a := C.App{}
	a.Path = *path
	a.Port = *port
	a.Epoch = epoch
	return &a
}

func main() {
	flag.Parse()
	app := newApp()


	// http.HandleFunc("/", handler)
	// http.HandleFunc("/health", C.HealthHandler)
	// if err := http.ListenAndServe(":" + strconv.Itoa(*port), nil); err != nil {
	// 	panic(err)
	// }
	
}