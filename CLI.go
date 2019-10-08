package poker

import (
	"io"
)

type CLI struct {
	playerStore PlayerStore
	input       io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Christopher")
}
