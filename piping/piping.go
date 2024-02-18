package piping

import (
	"bufio"
	"io"
	"strings"
)

type Pipeline struct {
	Data   io.Reader
	Output io.Writer
	Error  error
}

func FromString(data string) *Pipeline {
	p := Pipeline{
		Data: strings.NewReader(data),
	}

	return &p
}

func (p *Pipeline) Stdout() {
	input := bufio.NewScanner(p.Data)
	for input.Scan() {
		_, _ = p.Output.Write(input.Bytes())
	}
}
