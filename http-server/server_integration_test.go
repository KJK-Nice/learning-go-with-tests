package poker_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	poker "github.com/KJK-Nice/learning-go-with-tests/http-server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := poker.NewInMemoryPlayerStore()
	server := mustMakePlayerServer(t, store, dummyGame)
	player := "Pepper"

	t.Run("it run record win correctly", func(t *testing.T) {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response, http.StatusOK)

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
		assertStatus(t, response, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "500")

	})
}

func TestRecordingWinsAndRetrivingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := mustMakePlayerServer(t, store, dummyGame)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})

}
