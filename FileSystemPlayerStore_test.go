package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	initialData := `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Christopher", "Wins": 33}]`

	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

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
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Christopher")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("stores wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Christopher")

		got := store.GetPlayerScore("Christopher")
		want := 34

		assertScoreEquals(t, got, want)
	})

	t.Run("stores win for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1

		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tempFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create tmp file %v", err)
	}

	_, _ = tempFile.Write([]byte(initialData))

	removeFile := func() {
		_ = tempFile.Close()
		_ = os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}

func assertScoreEquals(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
