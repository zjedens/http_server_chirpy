package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.Handle("/app/assets/", http.StripPrefix("/app/assets/", http.FileServer(http.Dir("assets"))))
	mux.HandleFunc(
		"/healthz",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 OK"))
		},
	)

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Panic(server.ListenAndServe())
}
