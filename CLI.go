package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "

// CLI interface for the poker application
type CLI struct {
	playerStore PlayerStore
	input       *bufio.Scanner
	output      io.Writer
	game        *Game
}

// NewCLI creates a new CLI instance
func NewCLI(store PlayerStore, input io.Reader, output io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		input:       bufio.NewScanner(input),
		output:      output,
		game:        &Game{alerter, store},
	}
}

func (cli *CLI) PlayPoker() {
	_, _ = fmt.Fprint(cli.output, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.game.Finish(winner)
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.input.Scan()

	return cli.input.Text()
}
