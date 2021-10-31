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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Add("Expires", epoch)
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("X-Accel-Expires", "0")

	http.FileServer(http.Dir(*path)).ServeHTTP(w, r)
}


func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":" + strconv.Itoa(*port), nil); err != nil {
		panic(err)
	}
	
}