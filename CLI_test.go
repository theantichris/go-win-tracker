package poker_test

import (
	poker "github.com/theantichris/go-win-tracker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record christopher win from user input", func(t *testing.T) {
		input := strings.NewReader("Christopher wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, input)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Christopher")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		input := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, input)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
