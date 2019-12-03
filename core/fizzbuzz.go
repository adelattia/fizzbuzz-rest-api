package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Response struct {
	StatusCode uint
	Result     []string
	Message    string
}

// FizzBuzz Handler that accepts five params
// 3 ints : int1, int2, limit and 2 strings: string1 and string2
// all multiples of int1 are replaced by string1,
// all multiples of int2 are replaced by string2,
// all multiples of int1 and int2 are replaced by string1string2
func FizzBuzz(w http.ResponseWriter, r *http.Request) {

	int1, err := strconv.Atoi(r.FormValue("int1"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, errors.New("error parsing int1"), 1)
		return
	}
	int2, err := strconv.Atoi(r.FormValue("int2"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, errors.New("error parsing int2"), 1)
		return
	}
	string1 := r.FormValue("string1")
	string2 := r.FormValue("string2")

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, errors.New("error parsing limit"), 1)
		return
	}

	fizzBuzzList, error, statusCode := FizzBuzzNaive(int1, int2, string1, string2, limit)
	if error != nil {
		respondWithError(w, http.StatusBadRequest, error, statusCode)
		return
	}
	respondWithJSON(w, http.StatusOK, Response{StatusCode: statusCode, Result: fizzBuzzList, Message: ""})
}

func fizzBuzzCheck(int1, int2 int, string1, string2 string, limit int) ([]string, error, uint) {
	// check if int1 and int2 are the same
	if int1 == int2 {
		return []string{}, errors.New("You should provide different multiples"), 1
	}
	if int1 > limit || int2 > limit {
		return []string{}, errors.New("Multiples should be less than limit"), 1
	}
	return []string{}, nil, 0
}

// FizzBuzzNaive is a straightforward implementation using the modulo
func FizzBuzzNaive(int1, int2 int, string1, string2 string, limit int) ([]string, error, uint) {
	var result []string
	result, err, code := fizzBuzzCheck(int1, int2, string1, string2, limit)
	if err != nil {
		return result, err, code
	}
	for i := 1; i <= limit; i++ {
		if i%int1 == 0 && i%int2 == 0 {
			result = append(result, string1+string2)
		} else if i%int1 == 0 {
			result = append(result, string1)
		} else if i%int2 == 0 {
			result = append(result, string2)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result, nil, 0
}

// FizzBuzzNoModulo is based on an algorithm I found in https://news.ycombinator.com/item?id=1383978
// the purpose is to see if the modulo is a costly operation
func FizzBuzzNoModulo(int1, int2 int, string1, string2 string, limit int) ([]string, error, uint) {
	var result []string
	result, err, code := fizzBuzzCheck(int1, int2, string1, string2, limit)
	if err != nil {
		return result, err, code
	}
	i1 := 0
	i2 := 0
	for i := 1; i <= limit; i++ {
		i1++
		i2++
		if i1 == int1 && i2 == int2 {
			result = append(result, string1+string2)
		} else if i1 == int1 {
			result = append(result, string1)
			i1 = 0
		} else if i2 == int2 {
			result = append(result, string2)
			i2 = 0
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result, nil, 0
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, error error, statusCode uint) {

	errorResponse := Response{
		StatusCode: statusCode,
		Message:    error.Error(),
		Result:     nil,
	}

	response, _ := json.Marshal(errorResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
