package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pkgData struct {
	Name      string `json:"name"`
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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return packages, err
	}

	err = json.Unmarshal(data, &packages)
	if err != nil {
		return packages, err
	}
	return packages, nil
}

// getMovieHandler marshal movie data
func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	url := "http://go_api_db:8080/search" // change server url
	packages, err := patchData(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	jsonData, _ := json.Marshal(packages)

	fmt.Fprintf(w, string(jsonData))
}
