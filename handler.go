package optimizer

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	url := query.Get("url")
	wi := query.Get("w")
	h := query.Get("h")
	q := query.Get("q")

	fmt.Println(url, wi, h, q)
	accelRedirect := fmt.Sprintf("%s", url)
	w.Header().Set("X-Accel-Redirect", accelRedirect)
}
