package main

import (
	"flag"
	"net/http"
	"strconv"
	"time"
	C "./common"
	M "./middleware"
	Ctrl "./controllers"
)

// params
var (
	port = flag.Int("port", 8000, "port number")
	path = flag.String("path", ".", "path")
	epoch = time.Unix(0, 0).Format(time.RFC1123)
)


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


	http.HandleFunc("/", M.MiddlewareChain(Ctrl.FileHttpServer, *app))
	http.HandleFunc("/health", M.MiddlewareChain(C.HealthHandler, *app))
	if err := http.ListenAndServe(":" + strconv.Itoa(app.Port), nil); err != nil {
		panic(err)
	}
	
}