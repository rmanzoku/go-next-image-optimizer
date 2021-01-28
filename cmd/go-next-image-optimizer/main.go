package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"

	optimizer "github.com/rmanzoku/go-next-image-optimizer"
)

var (
	version  = "none"
	revision = "none"
	port     = "9900"
	imageSrc = ""
)

func init() {
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("IMAGE_SRC") != "" {
		imageSrc = os.Getenv("IMAGE_SRC")
	} else {
		panic("IMAGE_SRC environment value is not found")
	}
}

func main() {
	o := optimizer.NewOptimizer(imageSrc)
	http.HandleFunc("/", o.Handler)
	http.ListenAndServe(":"+port, nil)
}
