package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrorMessage = "Bad value received for number of players, please try again with a number"

// CLI interface for the poker application
type CLI struct {
	input  *bufio.Scanner
	output io.Writer
	game   Game
}

// NewCLI creates a new CLI instance
func NewCLI(input io.Reader, output io.Writer, game Game) *CLI {
	return &CLI{bufio.NewScanner(input), output, game}
}

func (cli *CLI) PlayPoker() {
	_, _ = fmt.Fprint(cli.output, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		_, _ = fmt.Fprint(cli.output, BadPlayerInputErrorMessage)

		return
	}

	cli.game.Start(numberOfPlayers, cli.output)

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
