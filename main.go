package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found page 404", 404)
	})
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
