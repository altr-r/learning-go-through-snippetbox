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
	w.Write([]byte("Hello from SnippetBox")) // require raw bytes, thus the conversion from string to bytes
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippets...."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		/*
		* .Set() replaces if there is an existing header
		* .Add() appends to existing header
		* .Del() delets all values of the given header
		* .Get() retrieves the first value
		* .Values() gets a slice of all
		 */
		w.Header().Set("Allow", "POST")

		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		/*
		* This above code snippet can be replaced by the line written below
		* which is a shortcut that calls w.WriteHeader() and w.Write() underneath
		* In terms of functionality it is exactly the same
		* The main difference is now http.ResponseWriter is being passed to another function
		* Which sends the response to the user for us
		 */
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet...."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
