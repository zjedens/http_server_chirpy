package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type chirp struct {
	Body string `json: body`
}

type validationResponse struct {
	Valid bool   `json: valid`
	Error string `json: error`
}

func addInternalServerError(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"error\":\"Internal server error\"}"))
}

func ValidateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	newChirp := chirp{}
	err := decoder.Decode(&newChirp)
	if err != nil {
		log.Printf("Request error: cannot deserialize")
		addInternalServerError(w)
		return
	}

	log.Printf("Received chirp:  %v", newChirp.Body)

	if len(newChirp.Body) > 140 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\":\"Chirp exceeds max character length\"}"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{\"valid\":true}"))
}
