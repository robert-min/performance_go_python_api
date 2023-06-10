package handlers

import (
	"fmt"
	"net/http"
)

type pkgData struct {
	UserId    string `json:"name"`
	MovieId   string `json:"movieId"`
	Rating    string `json:"rating"`
	TimeStamp string `json:"timestamp"`
}

// getMovieHandler marshal movie data
func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "hello")
}
