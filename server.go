package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"text/template"
)

const contentTypeHeader = "content-type"
const jsonContentType = "application/json"
const htmlTemplatePath = "game.html"

// Player stores a name with a number of wins
type Player struct {
	Name string
	Wins int
}

// PlayerStore stores score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
}

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore) (*PlayerServer, error) {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles("game.html")

	if err != nil {
		return nil, fmt.Errorf("problem loading template %s %v", htmlTemplatePath, err.Error())
	}

	p.template = tmpl
	p.store = store

	router := http.NewServeMux()
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/ws", http.HandlerFunc(p.websocketHandler))

	p.Handler = router

	return p, nil
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (p *PlayerServer) websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := wsUpgrader.Upgrade(w, r, nil)
	_, winnerMsg, _ := conn.ReadMessage()
	p.store.RecordWin(string(winnerMsg))
}

func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	_ = p.template.Execute(w, nil)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeHeader, jsonContentType)
	_ = json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	_, _ = fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
