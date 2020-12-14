package main

import (
	"fmt"
	"net/http"
	"os"
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
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	url := query.Get("url")
	wi := query.Get("w")
	h := query.Get("h")
	q := query.Get("q")

	fmt.Println(url, wi, h, q)
	accelRedirect := fmt.Sprintf("%s", url)
	w.Header().Set("X-Accel-Redirect", accelRedirect)
}
