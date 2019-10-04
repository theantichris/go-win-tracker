package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	_, _ = f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)

	return league
}
