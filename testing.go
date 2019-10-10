package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// StubPlayerStore is a stub used for testing PlayerStores
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

// GetPlayerScore gets the Player's score
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

// RecordWin records a win for the specified Player
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// GetLeague returns a list of Players and their scores
func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// NewGetScoreRequest returns a GET request for a Player's score
func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return req
}

// NewGetGameRequest returns a GET request for the /game endpoint
func NewGetGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/game"), nil)

	return req
}

// NewPostWinRequest returns a POST request to add a win to the Player
func NewPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)

	return request
}

// NewLeagueRequest returns a GET request to get the league
func NewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)

	return req
}

// GetLeagueFromResponse returns a league from a response body
func GetLeagueFromResponse(t *testing.T, body io.Reader) ([]Player, error) {
	t.Helper()

	league, err := NewLeague(body)

	return league, err
}

// AssertResponseBody asserts the body contains the expected data
func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// AssertStatus asserts the correct status was received
func AssertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

// AssertLeague asserts the correct data was returned for the league
func AssertLeague(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// AssertContentType asserts the content-type header has the wanted value
func AssertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	if response.Result().Header.Get(contentTypeHeader) != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
	}
}

// AssertPlayerWin asserts the Player has the correct amount of wins
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin but want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}
