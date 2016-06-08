package main

import (
	"net/http"
)

func main() {
	var shouldFail bool
	http.Handle("/check", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if shouldFail {
			http.Error(w, "fail", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("pass"))
	}))
	http.Handle("/fail", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shouldFail = true
	}))
	http.Handle("/pass", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shouldFail = false
	}))

	http.ListenAndServe(":8888", nil)
}
