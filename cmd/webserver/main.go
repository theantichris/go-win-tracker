package main

import (
	"log"
	"net/http"
	"os"

	poker "github.com/theantichris/go-win-tracker"
)

func main() {
	store := poker.NewInMemoryPlayerStore()
	server := poker.NewPlayerServer(store)

	port := getPort()
	if err := http.ListenAndServe(port, server); err != nil {
		log.Fatalf("could not listen on port %q, %v", port, err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	return ":" + port
}
