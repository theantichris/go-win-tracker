package poker_test

import (
	"bytes"
	"fmt"
	poker "github.com/theantichris/go-win-tracker"
	"io"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 7)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("butts\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrorMessage)
	})
}

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &StubPlayerStore{}

type GameSpy struct {
	StartCalled      bool
	StartCalledWith  int
	BlindAlert       []byte
	FinishCalled     bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int, alertsDestination io.Writer) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers

	_, _ = alertsDestination.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalledWith = winner
}

type ScheduledAlert struct {
	at     time.Duration
	amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, ScheduledAlert{duration, amount})
}

func assertScheduledAlert(t *testing.T, got, want ScheduledAlert) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t *testing.T, game *GameSpy, want int) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartCalledWith == want
	})

	if !passed {
		t.Errorf("wanted game to start with %d players got %d players", want, game.StartCalledWith)
	}
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishCalledWith == winner
	})

	if !passed {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishCalledWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)

	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}

	return false
}
