package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on port 4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View a snippet here..."))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a snippet..."))
}
