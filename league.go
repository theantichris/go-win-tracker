package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

// League is a slice of Players
type League []Player

// FindPlayer returns the specified Player from the League
func (l League) FindPlayer(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}

// NewLeague creates a new League instance
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player

	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("profile parsing league, %v", err)
	}

	return league, err
}
