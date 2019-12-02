package fizzbuzz

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// FizzBuzz Customized
func FizzBuzz(w http.ResponseWriter, r *http.Request) {

	int1, err := strconv.Atoi(r.FormValue("int1"))
	if err != nil {
		// todo: http error
	}
	int2, err := strconv.Atoi(r.FormValue("int2"))
	if err != nil {
		// todo: http error
	}
	string1 := r.FormValue("string1")
	string2 := r.FormValue("string2")

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		// todo: http error
	}

	var fizzBuzzList []string
	fizzBuzzNaive(int1, int2, string1, string2, limit, &fizzBuzzList)
	respondWithJSON(w, http.StatusOK, fizzBuzzList)
}

func fizzBuzzNaive(int1, int2 int, string1, string2 string, limit int, result *[]string) {
	for i := 1; i <= limit; i++ {
		if i%int1 == 0 && i%int2 == 0 {
			*result = append(*result, string1+string2)
		} else if i%int1 == 0 {
			*result = append(*result, string1)
		} else if i%int2 == 0 {
			*result = append(*result, string2)
		} else {
			*result = append(*result, strconv.Itoa(i))
		}
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
