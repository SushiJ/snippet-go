package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Specific snippet with id %d", id)
}
