package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Ryan", "Wins": 100}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Ryan", 100},
		}

		assertLeague(t, got, want)

		// We want to run again, making sure Seek from the io.ReadSeeker
		// jumps to the start of the file and reads it bit by bit `[]byte` again.
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}
