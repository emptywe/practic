package main

import (
	"net/http"
)

func checkSubStr(s string) string {
	str := []rune(s)
	appeared := make(map[rune]int)
	sequence := 0
	strn := []rune{}
	delta := 0
	for i := 0; i < len(str); i++ {

		if idx, ok := appeared[str[i]]; ok {
			delta = max(idx, delta)
		}
		sequence = max(sequence, i-delta+1)
		appeared[str[i]] = i + 1
		if len(strn) < len(str[delta:i+1]) {
			strn = str[delta : i+1]
		}
	}
	return string(strn)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	task := r.Header.Get("Task")
	if task == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Answer", checkSubStr(task))
}
