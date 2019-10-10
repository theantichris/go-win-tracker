package poker

import (
	"os"
)

type Tape struct {
	file *os.File
}

func NewTape(file *os.File) *Tape {
	return &Tape{file}
}

func (t *Tape) Write(p []byte) (n int, err error) {
	_ = t.file.Truncate(0)
	_, _ = t.file.Seek(0, 0)

	return t.file.Write(p)
}
