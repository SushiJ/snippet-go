package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost) //MethodPost - "POST"
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not allowed"))
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed) // StatusMethodNotAllowed = 405
		return
	}
	w.Write([]byte("CreateSnippet"))
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View Snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Live on http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
