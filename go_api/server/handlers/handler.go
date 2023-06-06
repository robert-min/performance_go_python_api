package handlers

import (
	"fmt"
	"net/http"
)

func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "movie handler")
}
