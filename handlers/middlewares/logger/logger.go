package logger

import (
	"fmt"
	"net/http"

	"github.com/SmurfsAtWork/lilpapa/log"
)

func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infoln(r.Method, r.URL.Path, r.URL.Query())
		for key, val := range r.Header {
			log.Infoln("HEADER ", key, val)
		}
		fmt.Println()
		h.ServeHTTP(w, r)
	})
}
