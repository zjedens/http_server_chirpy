package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
)

type chirp struct {
	Body string `json:"body"`
}

type validationResponse struct {
	Valid       bool   `json:"valid"`
	Error       string `json:"error"`
	CleanedBody string `json:"cleaned_body"`
}

func addJSONHeader(w http.ResponseWriter) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Add("Content-Type", "application/json")
	}
}

func addInternalServerError(w http.ResponseWriter) {
	addJSONHeader(w)
	w.WriteHeader(http.StatusInternalServerError)
	err := fmt.Sprintf("{\"error\":\"%v\"}", http.StatusText(http.StatusInternalServerError))
	w.Write([]byte(err))
}

func addValidationResponse(w http.ResponseWriter, status int, res validationResponse) {
	data, err := json.Marshal(res)

	if res.Valid == false {
		log.Printf("status: %v | error: %v", status, res.Error)
	}

	if err != nil {
		log.Printf("Response error: failed to serialize validation info")
		addInternalServerError(w)
		return
	}

	addJSONHeader(w)
	w.WriteHeader(status)
	w.Write(data)
}

func cleanBody(body string) string {
	badWords := []string{"kerfuffle", "sharbert", "fornax"}

	words := strings.Split(body, " ")

	for i := 0; i < len(words); i++ {
		if slices.Contains(badWords, strings.ToLower(words[i])) {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}

func ValidateChirp(w http.ResponseWriter, r *http.Request) {
	newChirp := chirp{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newChirp)
	if err != nil {
		log.Printf("Request error: cannot deserialize")
		addInternalServerError(w)
		return
	}

	log.Printf("Received chirp:  %v", newChirp.Body)

	if len(newChirp.Body) > 140 {
		addValidationResponse(w, http.StatusBadRequest, validationResponse{Valid: false, Error: "Chirp exceeds max character length"})
		return
	}

	addValidationResponse(
		w,
		http.StatusOK,
		validationResponse{Valid: true, CleanedBody: cleanBody(newChirp.Body)},
	)
}
