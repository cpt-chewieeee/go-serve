package main

import (
	"flag"
	"net/http"
	"strconv"
	"time"
)


var (
	port = flag.Int("port", 8000, "port number")
	path = flag.String("path", ".", "Path [default .]")
	epoch = time.Unix(0, 0).Format(time.RFC1123)
)

var defaultHeaders = map[string]string{
	"Access-Control-Allow-Origin": "*",
	"Expires": epoch,
	"Cache-Control": "no-cache, private, max-age=0",
	"Pragma": "no-cache",
	"X-Accel-Expires": "0",
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range defaultHeaders {
		w.Header().Add(k, v)
	}

	http.FileServer(http.Dir(*path)).ServeHTTP(w, r)
}


func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":" + strconv.Itoa(*port), nil); err != nil {
		panic(err)
	}
	
}