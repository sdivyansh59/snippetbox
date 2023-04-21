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
	w.Write([]byte("Hello from snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying  a specific snippet with ID %d ...", id)
	// w.Write([]byte("Displaying  a specific snippet witj ID"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		
		// w.Header().Set("Divyansh", "Singh")   added
		// w.WriteHeader(http.StatusMethodNotAllowed)  // 405
		// w.Write([]byte("Method not allowed"))
		// w.Header().Set("Divyansh", "Singh") not added

		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Crete a new snippet"))
}