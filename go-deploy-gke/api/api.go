package api

import (
	"flag"
	"go-deploy-gke/api/handlers"
	"go-deploy-gke/api/httpext"
	"go-deploy-gke/api/middleware"
	"log"
	"net/http"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8080, "set api port")
	flag.Parse()
}

func Run() {
	http.HandleFunc("/", middleware.Logger(handlers.Index))
	listen(port, func() {
		log.Printf("Listening [::]:%d...\n\n", port)
	})
}

func listen(port int, fns ...func()) {
	for _, fn := range fns {
		fn()
	}
	log.Fatalln(http.ListenAndServe(httpext.Port(port).Addr(), nil))
}
