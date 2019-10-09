package poker

import (
	"bufio"
	"io"
	"strings"
)

// CLI interface for the poker application
type CLI struct {
	playerStore PlayerStore
	input       *bufio.Scanner
}

// NewCLI creates a new CLI instance
func NewCLI(store PlayerStore, input io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		input:       bufio.NewScanner(input),
	}
}

// PlayPoker records a win for the user read from input
func (cli *CLI) PlayPoker() {
	input := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(input))
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.input.Scan()

	return cli.input.Text()
}
