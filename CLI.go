package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// CLI interface for the poker application
type CLI struct {
	playerStore PlayerStore
	input       *bufio.Scanner
	alerter     BlindAlerter
}

// NewCLI creates a new CLI instance
func NewCLI(store PlayerStore, input io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		input:       bufio.NewScanner(input),
		alerter:     alerter,
	}
}

// PlayPoker records a win for the user read from input
func (cli *CLI) PlayPoker() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}

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
