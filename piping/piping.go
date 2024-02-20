package piping

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
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

func (p *Pipeline) Freq() *Pipeline {
	type Counter struct {
		value string
		count int
	}

	if p.Error != nil {
		errorPip := FromString("")
		errorPip.Error = p.Error
		return errorPip
	}

	if p.Data == nil {
		errorPip := FromString("")
		errorPip.Error = errors.New("no data to filter.")
		return errorPip
	}

	input := bufio.NewScanner(p.Data)
	buf := new(bytes.Buffer)
	counter := Counter{}
	for input.Scan() {
		line := input.Text()
		if counter.value != line {
			if counter.count > 0 {
				fmt.Fprintf(buf, "%d %v\n", counter.count, counter.value)
			}
			counter = Counter{value: line, count: 0}
		}
		counter.count++
	}
	fmt.Fprintf(buf, "%d %v\n", counter.count, counter.value)

	res := New()
	res.Data = buf

	return res
}

func (p *Pipeline) Sort(descending bool) *Pipeline {
	if p.Error != nil {
		errorPip := FromString("")
		errorPip.Error = p.Error
		return errorPip
	}

	if p.Data == nil {
		errorPip := FromString("")
		errorPip.Error = errors.New("no data to sort.")
		return errorPip
	}
	content, err := p.String()
	if err != nil {
		errorPip := FromString("")
		errorPip.Error = errors.New("no data to sort.")
		return errorPip
	}
	lines := strings.Split(content, "\n")
	sort.Slice(lines, func(i, j int) bool {
		if descending {
			return lines[i] > lines[j]
		}
		return lines[i] < lines[j]
	})

	buf := new(bytes.Buffer)
	for _, el := range lines {
		if el == "" {
			continue
		}
		fmt.Fprintln(buf, el)
	}

	res := New()
	res.Data = buf

	return res
}

func (p *Pipeline) First(n int) *Pipeline {
	if p.Error != nil {
		errorPip := FromString("")
		errorPip.Error = p.Error
		return errorPip
	}

	if p.Data == nil {
		errorPip := FromString("")
		errorPip.Error = errors.New("no data to sort.")
		return errorPip
	}

	if n < 1 {
		errorPip := FromString("")
		errorPip.Error = errors.New("invalid argument")
		return errorPip
	}

	input := bufio.NewScanner(p.Data)
	buf := new(bytes.Buffer)
	for i := 0; input.Scan() && i < n; i++ {
		fmt.Fprintln(buf, input.Text())
	}
	res := New()
	res.Data = buf

	return res
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	io.Copy(p.Output, p.Data)
}
