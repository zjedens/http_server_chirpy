package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	config "github.com/zjedens/api/config"

	handler "github.com/zjedens/api/handler"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env config: %v", err)
	}
	dbUrl := os.Getenv(".env")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to backend: %v", err)
	}

	config := config.ApiConfig{}
	mux := http.NewServeMux()
	mux.Handle("/app/", config.MiddlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.Handle("/app/assets/", config.MiddlewareMetricsInc(http.StripPrefix("/app/assets/", http.FileServer(http.Dir("assets")))))
	mux.HandleFunc(
		"GET /api/healthz",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 OK\n"))
		},
	)
	mux.HandleFunc(
		"POST /api/validate_chirp",
		handler.ValidateChirp,
	)
	mux.HandleFunc("GET /admin/metrics", config.GetMiddlewareMetrics)
	mux.HandleFunc("POST /admin/reset", config.ResetMiddlewareMetrics)

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Panic(server.ListenAndServe())
}
