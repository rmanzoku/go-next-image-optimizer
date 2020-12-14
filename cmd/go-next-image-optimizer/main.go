package main

import (
	"net/http"
	"os"

	optimizer "github.com/rmanzoku/go-next-image-optimizer"
)

var (
	version  = "none"
	revision = "none"
	port     = "9900"
)

func init() {
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
}

func main() {
	http.HandleFunc("/", optimizer.Handler)
	http.ListenAndServe(":"+port, nil)
}
