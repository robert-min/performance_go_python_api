package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pkgData struct {
	UserId    string `json:"name"`
	MovieId   string `json:"movieId"`
	Rating    string `json:"rating"`
	TimeStamp string `json:"timestamp"`
}

// pathchData request to movie search server
func patchData(url string) ([]pkgData, error) {
	var packages []pkgData

	resp, err := http.Get(url)
	if err != nil {
		return packages, err
	}
	defer resp.Body.Close()

	if resp.Header.Get("Content-Type") != "application/json" {
		return packages, nil
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return packages, err
	}

	err = json.Unmarshal(data, &packages)
	return packages, err
}

// getMovieHandler marshal movie data
func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	url := "localhost:8080/search" // change server url
	packages, err := patchData(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	jsonData, err := json.Marshal(packages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintf(w, string(jsonData))
}
