package grep

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type finder struct {
	input  io.Reader
	output io.Writer
}

type option func(*finder) error

func WithInput(input io.Reader) option {
	return func(f *finder) error {
		if input == nil {
			return errors.New("nil reader")
		}
		f.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(f *finder) error {
		if output == nil {
			return errors.New("nil writer")
		}
		f.output = output
		return nil
	}
}

func NewFinder(opts ...option) (finder, error) {
	f := finder{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(&f)
		if err != nil {
			return finder{}, err
		}
	}

	return f, nil
}

func (f finder) Find(word string) []string {
	input := bufio.NewScanner(f.input)

	var foundStrings []string

	for input.Scan() {
		if curLine := input.Text(); strings.Contains(curLine, word) {
			foundStrings = append(foundStrings, curLine)
		}
	}

	return foundStrings
}

func Cmdline() {
	args := os.Args[1:]
	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err_f := NewFinder(WithInput(file))
	if err_f != nil {
		panic(err_f)
	}
	result := f.Find(args[1])

	for _, line := range result {
		fmt.Println(line)
	}
}
