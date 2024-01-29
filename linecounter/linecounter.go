package linecounter

import (
	"bufio"
	"io"
	"os"
)

type Counter struct {
	Input io.Reader
}

func NewCounter() Counter {
	counter := Counter{
		Input: os.Stdin,
	}

	return counter
}

func (c Counter) Count() int {
	lines := 0
	input := bufio.NewScanner(c.Input)

	for input.Scan() {
		lines++
	}

	return lines
}

func Main() int {
	return NewCounter().Count()
}
