package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) string {
	return "123"
}

func main() {
	// handler := http.HandlerFunc(PlayerServer)
	server := &Sub2{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

// インタフェース化、Sub2という名前のインターフェース
// func Sub2(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	fmt.Fprint(w, GetPlayerScore(player))
// }

type StubPlayerStore struct {
	scores map[string]string
}

func (s *StubPlayerStore) GetPlayerScore(name string) string {
	score := s.scores[name]
	return score
	// if name == "Pepper" {
	// 	return "20"
	// }

	// if name == "Floyd" {
	// 	return "10"
	// }

	// return ""
}

type PlayerStore interface {
	GetPlayerScore(name string) string
}

type Sub2 struct {
	store PlayerStore
}

func (p *Sub2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
