package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var TestsCheckSubStr = []struct {
	name     string
	task     string
	expected string
}{
	{"One letter test", "a", "a"},
	{"ABC test", "abc", "abc"},
	{"Five random letters test", "acdefacbde", "defacb"},
	{"Test Teso test", "testtesot", "teso"},
	{"Randomstring test", "randomstring", "andomstri"},
}

func Test_checkSubStr(t *testing.T) {
	for _, test := range TestsCheckSubStr {
		result := checkSubStr(test.task)
		if test.expected != result {
			t.Errorf("Failed: %s got: %s, expected: %s", test.name, result, test.expected)
		}
	}
}

var TestsMax = []struct {
	name     string
	a        int
	b        int
	expected int
}{
	{"One zero test", 1, 0, 1},
	{"Zero one test", 0, 1, 1},
	{"Zeros test", 0, 0, 0},
	{"Negative zero test", -1, 0, 0},
	{"Negative positive test", -1, 1, 1},
}

func Test_max(t *testing.T) {
	for _, test := range TestsMax {
		result := max(test.a, test.b)
		if result != test.expected {
			t.Errorf("Failed: %s, got: %d, expected: %d", test.name, result, test.expected)
		}
	}
}

var TestsServeSubstring = []struct {
	name               string
	method             string
	header             string
	expectedStatusCode int
}{
	{"OK test", "GET", "a", 200},
	{"Method not allowed test", "POST", "a", 405},
	{"Bad request test", "GET", "", 400},
}

func Test_serveSubstring(t *testing.T) {
	var h Handler
	for _, test := range TestsServeSubstring {
		rr, err := MakeRequest(test.method, "/api/substring", test.header, h.serveSubstring)
		if err != nil {
			t.Errorf("Can't make test request at %s:%s", test.name, err.Error())
		}
		if rr.Code != test.expectedStatusCode {
			t.Errorf("Failed %s, got: %d, expected: %d", test.name, rr.Code, test.expectedStatusCode)
		}
	}
}

func MakeRequest(method, url, header string, handlerFunc http.HandlerFunc) (*httptest.ResponseRecorder, error) {
	var (
		req = httptest.NewRequest(method, url, nil)
	)
	router := http.NewServeMux()
	router.HandleFunc(url, handlerFunc)

	req.Header.Set("Task", header)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr, nil
}
