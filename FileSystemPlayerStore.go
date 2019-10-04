package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	_, _ = f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)

	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().FindPlayer(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.FindPlayer(name)

	if player != nil {
		player.Wins++
	}

	_, _ = f.database.Seek(0, 0)
	_ = json.NewEncoder(f.database).Encode(league)
}
