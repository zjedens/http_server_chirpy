package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (config *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Calling middlewareMetricsInc(...)")
			config.fileserverHits.Add(1)
			fmt.Printf("New hits total: %v\n", config.fileserverHits.Load())

			next.ServeHTTP(w, r)
		},
	)
}

func (config *apiConfig) getMiddlewareMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling getMiddlewareMetrics(...)")

	count := int(config.fileserverHits.Load())
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits: " + strconv.Itoa(count)))
}

func (config *apiConfig) resetMiddlewareMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling resetMiddlewareMetrics(...)")

	config.fileserverHits.Store(0)

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset: " + strconv.Itoa(int(config.fileserverHits.Load()))))

}

func main() {
	config := apiConfig{}
	mux := http.NewServeMux()
	mux.Handle("/app/", config.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.Handle("/app/assets/", config.middlewareMetricsInc(http.StripPrefix("/app/assets/", http.FileServer(http.Dir("assets")))))
	mux.HandleFunc(
		"GET /api/healthz",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 OK\n"))
		},
	)
	mux.HandleFunc("GET /api/metrics", config.getMiddlewareMetrics)
	mux.HandleFunc("POST /api/reset", config.resetMiddlewareMetrics)

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Panic(server.ListenAndServe())
}
