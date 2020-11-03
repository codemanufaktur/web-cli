package cmd

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func serve() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/news", handleAll)
	r.HandleFunc("/news/{id}", handleSingle)

	r.HandleFunc("/api/news", handleAllApi).Methods("GET")
	r.HandleFunc("/api/news/{id}", handleSingleApi).Methods("GET")

	//TODO understand this 'go ...'
	go log.Fatal(http.ListenAndServe(":8080", r))
}

func handleAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Alle News</h1>")
}

func handleSingle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Nur eine News</h1>")
}

func handleAllApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Alle News</h1>")
}

func handleSingleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Nur eine News</h1>")
}