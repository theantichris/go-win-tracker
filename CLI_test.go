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

	assertPlayerWin(t, playerStore, "Christopher")
}
