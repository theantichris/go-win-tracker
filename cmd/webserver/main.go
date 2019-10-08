package main

import (
	"github.com/theantichris/go-win-tracker"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFile, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer closeFile()

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
