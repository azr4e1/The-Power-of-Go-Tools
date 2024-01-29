package linecounter

import (
	"bufio"
	"io"
	"os"
)

type counter struct {
	input  io.Reader
	output io.Writer
}

type option func(*counter)

func WithInput(input io.Reader) option {
	return func(c *counter) {
		c.input = input
	}
}

func NewCounter(opts ...option) counter {
	c := counter{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		opt(&c)
	}

	return c
}

func (c counter) Count() int {
	lines := 0
	input := bufio.NewScanner(c.input)

	for input.Scan() {
		lines++
	}

	return lines
}

func Main() int {
	return NewCounter().Count()
}
