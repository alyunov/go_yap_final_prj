package server

import (
	"fmt"
	"net/http"

	"github.com/alyunov/go_yap_final_prj/pkg/api"
)

func Run() error {
	port := 7540
	webDir := "./web"
	api.Init()
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
