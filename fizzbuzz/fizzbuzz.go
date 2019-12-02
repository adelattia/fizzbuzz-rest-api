package fizzbuzz

import (
	"encoding/json"
	"net/http"
)

func FizzBuzz(w http.ResponseWriter, r *http.Request) {

	fizzBuzzList := []string{}

	respondWithJSON(w, http.StatusOK, fizzBuzzList)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
