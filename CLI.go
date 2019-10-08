package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	input       io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.input)
	reader.Scan()

	cli.playerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
