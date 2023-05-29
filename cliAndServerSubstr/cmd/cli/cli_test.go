package main

import (
	"net/http/httptest"
	"testing"
)

var TestsSendRequest = []struct {
	name     string
	task     string
	url      string
	isErr    bool
	expected string
}{
	{"OK test", "abc", "", false, "abc"},
	{"Wrong url test", "", "abc", true, ""},
	{"Wrong task test", "", "", true, ""},
}

func Test_sendRequest(t *testing.T) {
	var (
		h   Handler
		url string
	)
	srv := httptest.NewServer(h)
	for _, test := range TestsSendRequest {
		if test.url == "" {
			url = srv.URL
		} else {
			url = test.url
		}
		answer, err := sendRequest(test.task, url)
		if err != nil && !test.isErr {
			t.Errorf("Failed: %s, unexpected err: %s", test.name, err.Error())
		}
		if answer != test.expected && !test.isErr {
			t.Errorf("Failed: %s, got: %s, expected: %s", test.name, answer, test.expected)
		}
	}
}
