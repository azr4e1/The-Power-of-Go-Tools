package main

import (
	"fmt"
	"io"
	"os"
)

type SafeWriter struct {
	Writer io.Writer
	Error  error
}

func (sw *SafeWriter) Write(data []byte) {
	if sw.Error != nil {
		return
	}
	_, err := sw.Writer.Write(data)
	sw.Error = err
}

func write(w io.Writer) error {
	metadata := []byte("hello\n")
	sw := SafeWriter{
		Writer: w,
	}

	sw.Write(metadata)
	sw.Write(metadata)
	sw.Write(metadata)
	sw.Write(metadata)

	return sw.Error
}

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = write(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
