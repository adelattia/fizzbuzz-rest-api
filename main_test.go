package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := a.initializeRoutes()
	router.ServeHTTP(rr, req)

	return rr
}

func TestMain(t *testing.T) {
	req, _ := http.NewRequest("GET", "/fizzbuzz?string1=Fizz&string2=Buzz&int1=3&int2=5&limit=15", nil)
	r := executeRequest(req)

	var got []string
	err := json.Unmarshal([]byte(r.Body.String()), &got)

	if err != nil {
		t.Errorf("Error processing the request")
		t.Fail()
	}

	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if !testEq(want, got) {
		t.Errorf("Invalid output\n was: %s\n expected: %s", got, want)
		t.Fail()
	}
}

func testEq(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
