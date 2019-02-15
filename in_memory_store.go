package main

// NewInMemoryPlayerStore returns a new instance of the InMemoryPlayerStore struct
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore is a temp, in memory, data store to use without a DB in dev
type InMemoryPlayerStore struct {
	store map[string]int
}

// RecordWin stub
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore stub
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
