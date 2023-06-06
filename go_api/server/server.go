package main

import (
	"log"
	"net/http"
	"os"

	"github.com/robert-min/performance_go_python_api/go_api/server/handlers"
)

func setupServer(mux *http.ServeMux) http.Handler {
	handlers.Register(mux)
	return mux
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		// Set go_api's port 8080
		listenAddr = ":8080"
	}

	mux := http.NewServeMux()

	wrappedMux := setupServer(mux)
	log.Fatal(http.ListenAndServe(listenAddr, wrappedMux))
}
