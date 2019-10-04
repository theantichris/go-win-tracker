package main

import (
	"io"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database := newDatabase()
		store := FileSystemPlayerStore{database}

		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Christopher", 33},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := newDatabase()
		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Christopher")
		want := 33

		assertScoreEquals(t, got, want)
	})
}

func newDatabase() io.ReadSeeker {
	database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Christopher", "Wins": 33}
		]`)

	return database
}

func assertScoreEquals(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
