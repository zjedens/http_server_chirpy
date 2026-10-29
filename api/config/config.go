package api

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3/database"
)

type ApiConfig struct {
	fileserverHits atomic.Int32
	queries        *database.Queries
}

func (config *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Calling middlewareMetricsInc(...)")
			config.fileserverHits.Add(1)
			fmt.Printf("New hits total: %v\n", config.fileserverHits.Load())

			next.ServeHTTP(w, r)
		},
	)
}

func (config *ApiConfig) GetMiddlewareMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling getMiddlewareMetrics(...)")

	count := int(config.fileserverHits.Load())
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	metricsTemplate := `
<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>
`
	w.Write([]byte(fmt.Sprintf(metricsTemplate, count)))
}

func (config *ApiConfig) ResetMiddlewareMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling resetMiddlewareMetrics(...)")

	config.fileserverHits.Store(0)

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset: " + strconv.Itoa(int(config.fileserverHits.Load()))))

}
