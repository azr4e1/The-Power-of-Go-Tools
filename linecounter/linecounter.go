package linecounter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input  io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output reader")
		}
		c.output = output
		return nil
	}
}

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return counter{}, err
		}
	}

	return c, nil
}

func (c counter) Count() int {
	lines := 0
	input := bufio.NewScanner(c.input)

	for input.Scan() {
		lines++
	}

	return lines
}

func Main() {
	c, _ := NewCounter()
	count := c.Count()
	fmt.Println("---------------\nNumber of lines:", count)
}
