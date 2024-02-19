package piping

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
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

func (p *Pipeline) String() (string, error) {
	if p.Error != nil {
		return "", p.Error
	}

	data, err := io.ReadAll(p.Data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (p *Pipeline) Column(n int) *Pipeline {
	if p.Error != nil {
		errorPip := FromString("")
		errorPip.Error = p.Error
		return errorPip
	}

	if n <= 0 {
		errorPip := FromString("")
		errorPip.Error = errors.New("invalid column.")
		return errorPip
	}

	if p.Data == nil {
		errorPip := FromString("")
		errorPip.Error = errors.New("no data to filter.")
		return errorPip
	}

	input := bufio.NewScanner(p.Data)
	buf := new(bytes.Buffer)
	for input.Scan() {
		line := input.Text()
		fields := strings.Fields(line)
		if len(fields) < n-1 {
			errorPip := FromString("")
			errorPip.Error = errors.New("not enough columns to filter.")
			return errorPip
		}

		fmt.Fprintln(buf, fields[n-1])
	}

	outPip := New()
	outPip.Data = buf

	return outPip
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	io.Copy(p.Output, p.Data)
}
