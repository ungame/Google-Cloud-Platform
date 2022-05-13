package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var (
			start         = time.Now()
			method        = request.Method
			host          = request.Host
			uri           = request.RequestURI
			contentLength = request.ContentLength
		)

		next(writer, request)

		log.Printf("%s %s%s %d %s", method, host, uri, contentLength, time.Since(start).String())
	}
}
