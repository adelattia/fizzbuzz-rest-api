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
	FizzBuzzNaive(int1, int2, string1, string2, limit, &fizzBuzzList)
	respondWithJSON(w, http.StatusOK, fizzBuzzList)
}

// FizzBuzzNaive is a straightforward implementation using the modulo
func FizzBuzzNaive(int1, int2 int, string1, string2 string, limit int, result *[]string) {
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

// FizzBuzzNoModulo is based on an algorithm I found in https://news.ycombinator.com/item?id=1383978
// the purpose is to see if the modulo is a costly operation
func FizzBuzzNoModulo(int1, int2 int, string1, string2 string, limit int, result *[]string) {
	i1 := 0
	i2 := 0
	for i := 1; i <= limit; i++ {
		i1++
		i2++
		if i1 == int1 && i2 == int2 {
			*result = append(*result, string1+string2)
		} else if i1 == int1 {
			*result = append(*result, string1)
			i1 = 0
		} else if i2 == int2 {
			*result = append(*result, string2)
			i2 = 0
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
