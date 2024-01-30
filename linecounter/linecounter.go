package linecounter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type counter struct {
	files  []io.Reader
	input  io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) < 1 {
			return nil
		}
		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c.files[i] = f
		}
		c.input = io.MultiReader(c.files...)
		return nil
	}
}

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
			return errors.New("nil output writer")
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

func (c counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.input)

	for input.Scan() {
		lines++
	}

	for _, f := range c.files {
		f.(io.Closer).Close()
	}

	return lines
}

func (c counter) Words() int {
	wordsNr := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		line := input.Text()
		words := strings.Split(line, " ")
		for _, word := range words {
			if strings.TrimSpace(word) == "" {
				continue
			}
			wordsNr++
		}
	}
	return wordsNr
}

func MainLines() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	count := c.Lines()
	fmt.Println("Number of lines:", count)
	return 0
}

func MainWords() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	count := c.Words()
	fmt.Println("Number of lines:", count)
	return 0
}
