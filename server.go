package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	if player == "Pepper" {
		_, _ = fmt.Fprint(w, "20")
	}

	if player == "Floyd" {
		_, _ = fmt.Fprint(w, "10")
	}
}
