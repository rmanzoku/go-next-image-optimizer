package optimizer

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (o *Optimizer) Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	url := query.Get("url")
	if url == "" {
		http.Error(w, "Not Found", 404)
		return
	}

	var err error

	var wi int
	if query.Get("w") != "" {
		wi, err = strconv.Atoi(query.Get("w"))
		if err != nil {
			http.Error(w, fmt.Sprintf("...: %w", err), 400)
			return
		}
	}
	var h int
	if query.Get("h") != "" {
		h, err = strconv.Atoi(query.Get("h"))
		if err != nil {
			http.Error(w, fmt.Sprintf("...: %w", err), 400)
			return
		}
	}
	var q int = 100
	if query.Get("q") != "" {
		q, err = strconv.Atoi(query.Get("q"))
		if err != nil {
			http.Error(w, fmt.Sprintf("...: %w", err), 400)
			return
		}
	}

	accept := r.Header.Get("Accept")
	webpReady := strings.Contains(accept, "image/webp")

	ret, err := o.Optimize(url, wi, h, q, webpReady)
	if err != nil {
		http.Error(w, fmt.Sprintf("...: %w", err), 500)
		return
	}

	contentType := "image/png"
	if webpReady {
		contentType = "image/webp"
	}
	w.Header().Set("Content-Type", contentType)
	w.Write(ret)
}
