package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetreivingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Ryan"

	server.ServeHTTP(httptest.NewRecorder(), newPostWithRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWithRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWithRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromRequest(t, response.Body)
		want := []Player{
			{"Ryan", 3},
		}
		assertLeague(t, got, want)
	})
}
