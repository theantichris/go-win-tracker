package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := NewPlayerServer(store)

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
