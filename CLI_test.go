package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record christopher win from user input", func(t *testing.T) {
		input := strings.NewReader("Christopher wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, input}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Christopher")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		input := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, input}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cleo")
	})
}
