package piping

import (
	"io"
	"os"
	"strings"
)

type Pipeline struct {
	Data   io.Reader
	Output io.Writer
	Error  error
}

func New() *Pipeline {
	return &Pipeline{
		Output: os.Stdout,
	}
}

func FromString(data string) *Pipeline {
	p := New()
	p.Data = strings.NewReader(data)

	return p
}

func FromFile(filename string) *Pipeline {
	p := New()
	reader, err := os.Open(filename)
	if err != nil {
		return &Pipeline{Error: err}
	}
	p.Data = reader

	return p
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	io.Copy(p.Output, p.Data)
}
