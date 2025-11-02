package main

import (
	"github.com/KJK-Nice/learning-go-with-tests/http-server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal("problem creating player server", err)
	}

	if err := http.ListenAndServe(":5001", server); err != nil {
		log.Fatalf("could no listen on port 5000 %v", err)
	}
}
