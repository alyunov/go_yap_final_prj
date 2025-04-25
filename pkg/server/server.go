package server

import (
	"fmt"
	"net/http"
)

func Run() error {
	port := 7540
	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
