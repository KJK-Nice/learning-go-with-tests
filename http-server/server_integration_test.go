package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	t.Run("it run record win correctly", func(t *testing.T) {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		player := "Tony"
		wantedCount := 500

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
				wg.Done()
			}()
		}
		wg.Wait()

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "500")

	})
}
