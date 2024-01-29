package hellopeople

import (
	"bufio"
	"fmt"
	"io"
)

func ReadFrom(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	result, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return result, nil
}

func PrintTo(w io.Writer, name string) {
	fmt.Fprintln(w, "Hello,", name)
}
