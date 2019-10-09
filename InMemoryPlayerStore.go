package poker

// InMemoryPlayerStore stores Player data in memory
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetPlayerScore returns the number of wins for the specified Player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin adds a win to the specified Player's score
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetLeague returns a list of Players and scores
func (i *InMemoryPlayerStore) GetLeague() (league League) {
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return
}

// NewInMemoryPlayerStore creates and returns a new InMemoryPlayerStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}
