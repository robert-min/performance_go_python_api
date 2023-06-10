package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/robert-min/performance_go_python_api/go_api_db/server/lib"
)

// getMovieHandler marshal movie data
func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	packages, err := lib.MysqlConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	data, _ := json.Marshal(packages)

	fmt.Fprintf(w, string(data))
}
