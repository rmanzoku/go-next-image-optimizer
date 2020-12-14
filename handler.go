package optimizer

import (
	"fmt"
	"net/http"
	"strconv"
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
		h, err = strconv.Atoi(query.Get("w"))
		if err != nil {
			http.Error(w, fmt.Sprintf("...: %w", err), 400)
			return
		}
	}
	var q int
	if query.Get("q") != "" {
		q, err = strconv.Atoi(query.Get("w"))
		if err != nil {
			http.Error(w, fmt.Sprintf("...: %w", err), 400)
			return
		}
	}
	ret, err := o.Optimize(url, wi, h, q, true)
	if err != nil {
		http.Error(w, fmt.Sprintf("...: %w", err), 500)
		return
	}

	w.Header().Set("Content-Type", "image/webp")
	w.Write(ret)
}
