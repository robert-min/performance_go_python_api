package handlers

import "net/http"

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/search", getMovieHandler)
}
