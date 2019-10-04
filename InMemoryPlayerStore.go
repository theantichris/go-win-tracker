package main

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() (league League) {
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}
