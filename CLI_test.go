package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	input := strings.NewReader("Christopher wins\n")
	playerStore := &StubPlayerStore{}

	cli := &CLI{playerStore, input}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := playerStore.winCalls[0]
	want := "Christopher"

	if got != want {
		t.Errorf("didn't record correct winner, got %q want %q", got, want)
	}
}
